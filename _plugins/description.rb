require 'nokogiri'

module Jekyll
  module Description

    class Generator < Jekyll::Generator
      def generate(site)
        site.pages.each do |page|
          converter = site.find_converter_instance(::Jekyll::Converters::Markdown)
          dom = Nokogiri::HTML.fragment(converter.convert(page.content))
          dom.xpath("script").remove
          page.data['content-filtered'] = dom.to_html

          desc = dom.xpath("p")

          if desc.nil? || desc.length() == 0
            next
          end

          desc[0].css('sup').each do |src|
            src.remove
          end

          page.data['description'] = desc[0].text
        end
      end
    end

  end 
end 