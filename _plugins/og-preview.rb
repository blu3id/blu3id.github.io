require 'fileutils'

module OGPreview
  class OGPreviewGenerator < Jekyll::Generator
    safe true

    def generate(site)
      @og_previews = []

      site.pages.each do |page|
        if page.data['title'].nil?
          next
        end

        preview_doc = Jekyll::Document.new(
          page.path,
          {site: site, collection: site.collections['og-preview']}
        )
  
        preview_doc.read

        if preview_doc.data['layout'] != "default"
          next
        end

        # Set the layout to preview
        preview_doc.data['layout'] = 'og-preview'
        unless page.data['permalink'].nil?
          preview_doc.data['permalink'] = '/og-preview' + page.data['permalink']
          @og_previews.append({
            "output" => '_site/assets/og-preview/' + page.data['permalink'].split('/').last + '.png',
            "url" => '_site/og-preview' + page.data['permalink'] + '.html', 
            "width" => 600,
            "height" => 300
          })
        end

        # Add document to the collection
        site.collections['og-preview'].docs << preview_doc
      end

      FileUtils.mkdir_p '_site/og-preview'
      File.open('_site/og-preview/_og_previews.yml', "w") { 
        |file| file.write @og_previews.to_yaml[4..-1]
      }
    end
  end
end

# add to _config.yaml
# collections:
#   og-preview:
#     output: true
