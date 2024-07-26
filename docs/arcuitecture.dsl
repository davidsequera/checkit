workspace {
    
    !identifiers hierarchical
    
    model {
        user = person "User"
    
        softwareSystem = softwareSystem "Enron Email App System" {

            webapp = container "Web Application" "A page fo the explorarion of the emails database" "Vuejs" "Web Client"
            

            emails_service = container "Email Service" "Go (Chi)"{
                tags "Service"
                get_emails = component "Search Emails"
            }


            zinsearch = container "Email Search Engine" {
                tags "emails" "searchEngine"
                search_emails = component "Search for the emails"
            }
            user -> webapp "Uses" "Browser HTTPS"
            webapp -> emails_service.get_emails "Uses" "HTTPS"
            emails_service.get_emails -> zinsearch.search_emails "Uses HTTPS"
        }
    }
    
    views {

        systemLandscape softwareSystem "SystemLandscape" {
            include *
            autolayout
        }
        container softwareSystem "Containers_All" {
            include *
            autolayout lr
        }
        theme default

        styles {
            element "Web Client" {
                shape WebBrowser
            }
            element "searchEngine" {
                shape cylinder
            }
            element "Service" {
                background #F08CA4
            }
            
        }

    }

}