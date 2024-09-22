require 'uri'
require 'nokogiri'
require 'set'
require 'json'

module Jekyll
  module ArchiveLinks

    class Generator < Jekyll::Generator
      def generate(site)
        @outlinks = JSON.parse(File.read('outlinks.json')) if File.file? 'outlinks.json'
        @outlinks ||= {}
        @outlinks_check = {}

        site.pages.each do |page|
          if page.data['title'].nil? || page.data['draft'] || page.data['unlisted'] || page.data['date'].nil?
            next
          end

          converter = site.find_converter_instance(::Jekyll::Converters::Markdown)
          urls = get_links(converter.convert(page.content)).uniq.sort
          broken_urls = get_broken_links(converter.convert(page.content)).uniq.sort
          stable_prefix = Set['https://dx.doi.org', 'https://en.wikipedia.org/']
          page.data['outlinks'] = urls
          page.data['outlinks-archive'] = urls.reject {|page| page.start_with?(*stable_prefix)}
          save_outlinks(page.data['outlinks-archive'], page.path, page.data['date'], broken_urls)
        end

        File.open('outlinks.json', "w") { 
          |file| file.write JSON.pretty_generate(@outlinks.sort.to_h, :indent => "  ")
        }
        File.open('outlinks-check.json', "w") { 
          |file| file.write JSON.pretty_generate(@outlinks_check.sort.to_h, :indent => "  ")
        }
      end

      def get_links(content)
        Nokogiri::HTML(content).css("a").map do |link|
          if (href = link.attr("href")) && href.match(/^https?:/)
            uri = URI.parse(href)
            uri.fragment = nil
            href = uri.to_s
          end
        end.compact
      end

      def get_broken_links(content)
        Nokogiri::HTML(content).css("a.broken").map do |link|
          if (href = link.attr("href")) && href.match(/^https?:/)
            uri = URI.parse(href)
            uri.fragment = nil
            href = uri.to_s
          end
        end.compact
      end

      def save_outlinks(urls, permalink, date, broken_urls)
        stable_prefix = Set['https://webarchive.nationalarchives.gov.uk/',
          'https://web.archive.org/web/']
        urls.each do |url|
          if url.start_with?(*stable_prefix)
            m = url.match(%r#([1-2]\d{3}\d*.?)/(.*)#)
            fullurl = url
            capture = m[1]
            url = m[2]
          end

          if @outlinks.has_key? permalink
            unless @outlinks[permalink].has_key? url
              @outlinks[permalink][url] = date
            end
          else
            @outlinks[permalink] = {}
            @outlinks[permalink][url] = date
          end

          unless capture.nil?
            @outlinks[permalink][url] = fullurl
            next
          end

          if broken_urls.include? url
            next
          end

          if @outlinks_check.has_key? permalink
            unless @outlinks_check[permalink].include? url
              @outlinks_check[permalink] << url
            end
          else
            @outlinks_check[permalink] = [url]
          end
        end

        unless @outlinks[permalink].nil?
          @outlinks[permalink] = @outlinks[permalink].sort.to_h
        end
        unless @outlinks_check[permalink].nil?
          @outlinks_check[permalink] = @outlinks_check[permalink].sort
        end
      end

    end
  end 
end 