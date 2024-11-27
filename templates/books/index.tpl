{{ define "books/index.tpl"}}
<!DOCTYPE html>
    <html lang="en">
        <head>
            <meta charset="utf-8">
            <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
            <meta name="description" content="">
            <meta name="author" content="">

            <title>Blogs App</title>

            <!-- Custom fonts for this template -->
            <link href="https://fonts.googleapis.com/css?family=Lato:300,400,700,300italic,400italic,700italic" rel="stylesheet" type="text/css">

        </head>
    <body>
        <div>
            {{if .page}}
        <script>
            function generatePaginationLinks(link, currentPage, totalPages, pageSize) {
                let paginationHTML = '';

                if (totalPages > 1) {
                    paginationHTML += '<div class="pagination">';

                    if (currentPage > 1) {
                        paginationHTML += `<a href="${link}?page=${currentPage - 1}&pageSize=${pageSize}">Previous</a>`;
                    }

                    for (let page = 1; page <= totalPages; page++) {
                        if (page === currentPage) {
                            paginationHTML += `<span class="current-page">${page}</span>`;
                        } else {
                            paginationHTML += `<a href="${link}?page=${page}&pageSize=${pageSize}">${page}</a>`;
                        }
                    }

                    if (currentPage < totalPages) {
                        paginationHTML += `<a href="${link}?page=${currentPage + 1}&pageSize=${pageSize}">Next</a>`;
                    }

                    paginationHTML += '</div>';
                }

                return paginationHTML;
            }

            window.onload = function() {
                const paginationLinks = generatePaginationLinks("/books", {{.page}}, {{.totalPages}}, {{.pageSize}});

                const contentDiv = document.getElementById('pagination');
                contentDiv.innerHTML = paginationLinks;
            };
        </script>
        {{end}}

        <div>
            <h1>Books</h1>

            <div id="pagination" class="pagination">
            </div>

            <br/>
            <br/>

            <ul>
                {{ with .books }}
                    {{ range . }}
                        <li>
                            <div>
                                <a href="/books/{{.ID}}">
                                    <h5>{{ .Title }} </h5>
                                </a>
                            </div>
                            <br/>
                        </li>
                    {{ end }}
                {{ end }}
            </ul>
        </div>
        </div>
    </body>
{{end}}