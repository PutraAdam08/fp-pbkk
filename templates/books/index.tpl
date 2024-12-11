{{ define "books/index.tpl"}}
<!DOCTYPE html>
    <html lang="en">
        <head>
            <meta charset="utf-8">
            <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
            <meta name="description" content="">
            <meta name="author" content="">

            <title>Books App</title>

            <!-- Custom fonts for this template -->
            <link href="https://fonts.googleapis.com/css?family=Lato:300,400,700,300italic,400italic,700italic" rel="stylesheet" type="text/css">
            <link href="https://cdn.jsdelivr.net/npm/flowbite@2.5.2/dist/flowbite.min.css"  rel="stylesheet" />
            <script src="https://cdn.tailwindcss.com"></script>
        </head>
    <body class="bg-gray-50 dark:bg-gray-900 dark">
        <div>
            {{if .page}}
        <script>
            function generatePaginationLinks(link, currentPage, totalPages, pageSize) {
                let paginationHTML = '';

                if (totalPages > 1) {
                    paginationHTML += '<nav class="mt-10 flex items-center justify-center sm:mt-8" aria-label="Page navigation">';

                    paginationHTML += '<ul class="flex h-8 items-center -space-x-px text-sm">';

                    if (currentPage > 1) {
                        paginationHTML += `<li>
                        <a href="${link}?page=${currentPage - 1}&pageSize=${pageSize}" class="ms-0 flex h-8 items-center justify-center rounded-s-lg border border-e-0 border-gray-300 bg-white px-3 leading-tight text-gray-500 hover:bg-gray-100 hover:text-gray-700 dark:border-gray-700 dark:bg-gray-800 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white">
                        <span class="sr-only">Previous</span>
                        <svg class="h-4 w-4 rtl:rotate-180" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="none" viewBox="0 0 24 24">
                            <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15 19-7-7 7-7" />
                        </svg>
                        </a>
                        </li>`;
                    }
                    else if (currentPage <= 1) {
                        paginationHTML += `<li>
                        <span class="ms-0 flex h-8 items-center justify-center rounded-s-lg border border-e-0 border-gray-300 bg-white px-3 leading-tight text-gray-500 hover:bg-gray-100 hover:text-gray-700 dark:border-gray-700 dark:bg-gray-800 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white">
                        <span class="sr-only">Previous</span>
                        <svg class="h-4 w-4 rtl:rotate-180" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="none" viewBox="0 0 24 24">
                            <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15 19-7-7 7-7" />
                        </svg>
                        </span>
                        </li>`;
                    }

                    for (let page = 1; page <= totalPages; page++) {
                        if (page === currentPage) {
                            paginationHTML += `<li>
                                <span href="#" aria-current="page" class="z-10 flex h-8 items-center justify-center border border-primary-300 bg-primary-50 px-3 leading-tight text-primary-600 hover:bg-primary-100 hover:text-primary-700 dark:border-gray-700 dark:bg-gray-700 dark:text-white">${page}</span>
                            </li>`;
                        } else {
                            paginationHTML += `<li>
                            <a href="${link}?page=${page}&pageSize=${pageSize}" class="flex h-8 items-center justify-center border border-gray-300 bg-white px-3 leading-tight text-gray-500 hover:bg-gray-100 hover:text-gray-700 dark:border-gray-700 dark:bg-gray-800 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white">${page}</a>
                            </li>`;
                        }
                    }

                    if (currentPage < totalPages) {
                        paginationHTML += `<li>
                        <a href="${link}?page=${currentPage + 1}&pageSize=${pageSize}" class="flex h-8 items-center justify-center rounded-e-lg border border-gray-300 bg-white px-3 leading-tight text-gray-500 hover:bg-gray-100 hover:text-gray-700 dark:border-gray-700 dark:bg-gray-800 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white">
                        <span class="sr-only">Next</span>
                        <svg class="h-4 w-4 rtl:rotate-180" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="none" viewBox="0 0 24 24">
                            <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m9 5 7 7-7 7" />
                        </svg>
                        </a>
                        </li>`;
                    }
                    if (currentPage == totalPages) {
                        paginationHTML += `<li>
                        <span href="#" class="flex h-8 items-center justify-center rounded-e-lg border border-gray-300 bg-white px-3 leading-tight text-gray-500 hover:bg-gray-100 hover:text-gray-700 dark:border-gray-700 dark:bg-gray-800 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white">
                        <span class="sr-only">Next</span>
                        <svg class="h-4 w-4 rtl:rotate-180" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="none" viewBox="0 0 24 24">
                            <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m9 5 7 7-7 7" />
                        </svg>
                        </span>
                        </li>`;
                    }

                    paginationHTML += '</ul></nav>';
                }

                // if (totalPages > 1) {
                //     paginationHTML += '<div class="pagination">';

                //     if (currentPage > 1) {
                //         paginationHTML += `<a href="${link}?page=${currentPage - 1}&pageSize=${pageSize}">Previous</a>`;
                //     }

                //     for (let page = 1; page <= totalPages; page++) {
                //         if (page === currentPage) {
                //             paginationHTML += `<span class="current-page">${page}</span>`;
                //         } else {
                //             paginationHTML += `<a href="${link}?page=${page}&pageSize=${pageSize}">${page}</a>`;
                //         }
                //     }

                //     if (currentPage < totalPages) {
                //         paginationHTML += `<a href="${link}?page=${currentPage + 1}&pageSize=${pageSize}">Next</a>`;
                //     }

                //     paginationHTML += '</div>';
                // }

                return paginationHTML;
            }

            window.onload = function() {
                const paginationLinks = generatePaginationLinks("/books", {{.page}}, {{.totalPages}}, {{.pageSize}});

                const contentDiv = document.getElementById('pagination');
                contentDiv.innerHTML = paginationLinks;
            };
        </script>
        {{end}}

        <!-- <div>

            <ul>
                {{ with .books }}
                    {{ range . }}
                        <li>
                            <div>
                                <a href="/books/{{.ID}}">
                                    <h5>{{ .Title }} </h5>
                                </a>
                            </div> -->
                            <!-- <div class="mt-6 flow-root sm:mt-8">
                                <div class="divide-y divide-gray-200 dark:divide-gray-700">
                                <div class="flex flex-wrap items-center gap-y-4 py-6">
                                    <dl class="w-1/2 sm:w-1/4 lg:w-auto lg:flex-1">
                                    <dt class="text-base font-medium text-gray-500 dark:text-gray-400">Book Title:</dt>
                                    <dd class="mt-1.5 text-base font-semibold text-gray-900 dark:text-black">
                                        <a href="#" class="hover:underline">{{ .Title }} </a>
                                    </dd>
                                    </dl>
            
            
                                    <div class="w-full grid sm:grid-cols-2 lg:flex lg:w-64 lg:items-center lg:justify-end gap-4">
                                    <button type="button" class="w-full rounded-lg border border-red-700 px-3 py-2 text-center text-sm font-medium text-red-700 hover:bg-red-700 hover:text-white focus:outline-none focus:ring-4 focus:ring-red-300 dark:border-red-500 dark:text-red-500 dark:hover:bg-red-600 dark:hover:text-white dark:focus:ring-red-900 lg:w-auto">Cancel order</button>
                                    <a href="#" class="w-full inline-flex justify-center rounded-lg  border border-gray-200 bg-white px-3 py-2 text-sm font-medium text-gray-900 hover:bg-gray-100 hover:text-primary-700 focus:z-10 focus:outline-none focus:ring-4 focus:ring-gray-100 dark:border-gray-600 dark:bg-gray-800 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white dark:focus:ring-gray-700 lg:w-auto">View details</a>
                                    </div>
                                </div>
                            </div>
                            <br/>
                        </li>
                    {{ end }}
                {{ end }}
            </ul>
        </div> -->

        <section class="bg-gray-50 dark:bg-gray-900 py-8 antialiased md:py-16">

            <nav class="bg-white dark:bg-gray-900 fixed w-full z-20 top-0 start-0 border-b border-gray-200 dark:border-gray-600">
                <div class="max-w-screen-xl flex flex-wrap items-center justify-between mx-auto p-4">
                <a href="" class="flex items-center space-x-3 rtl:space-x-reverse">
                    <span class="self-center text-3xl font-semibold whitespace-nowrap dark:text-white">Library</span>
                </a>
                <div class="flex md:order-2 space-x-3 md:space-x-0 rtl:space-x-reverse">
                    <button type="button" onclick="logout()" class="text-red-600 border border-red-600 hover:text-white hover:bg-red-600 focus:ring-4 focus:outline-none focus:ring-red-600 font-medium rounded-lg text-bs px-10 py-2 text-center">Log Out</button>
                </button>
                </div>
                
                </div>
            </nav>

            <div class="mt-10 mx-auto max-w-screen-xl px-4 2xl:px-0">
                <div class="mx-auto max-w-5xl">
                <div class="gap-4 sm:flex sm:items-center sm:justify-between">
                    <h2 class="text-xl font-semibold text-gray-900 dark:text-white sm:text-3xl">Book List</h2>
                </div>

                <div class="mt-6 flow-root sm:mt-8">
                    <div class="divide-y divide-gray-200 dark:divide-gray-700">

                    {{ with .books }}
                        {{ range . }}
                        <div class="flex flex-wrap items-center gap-y-4 py-6">
                            <dl class="w-1/2 sm:w-1/4 lg:w-auto lg:flex-1">
                                <dt class="text-base font-medium text-gray-500 dark:text-gray-400">Book Title:</dt>
                                <dd class="mt-1.5 text-base font-semibold text-gray-900 dark:text-white">
                                    <a href="#" class="hover:underline">{{ .Title }}</a>
                                </dd>
                            </dl>

                            <dl class="w-1/2 sm:w-1/4 lg:w-auto lg:flex-1">
                                <dt class="text-base font-medium text-gray-500 dark:text-gray-400">Category:</dt>
                                <dd class="mt-1.5 text-base font-semibold text-gray-900 dark:text-white">
                                    <a href="#" class="hover:underline">{{ .Category }}</a>
                                </dd>
                            </dl>

                            <dl class="w-1/2 sm:w-1/4 lg:w-auto lg:flex-1">
                                <dt class="text-base font-medium text-gray-500 dark:text-gray-400">Published Year:</dt>
                                <dd class="mt-1.5 text-base font-semibold text-gray-900 dark:text-white">
                                    <a href="#" class="hover:underline">{{ .PublishYear }}</a>
                                </dd>
                            </dl>

                            <dl class="w-1/2 sm:w-1/4 lg:w-auto lg:flex-1">
                                <dt class="text-base font-medium text-gray-500 dark:text-gray-400">ISBN:</dt>
                                <dd class="mt-1.5 text-base font-semibold text-gray-900 dark:text-white">
                                    <a href="#" class="hover:underline">{{ .ISBN }}</a>
                                </dd>
                            </dl>
    
                        </div>
                        {{ end }}
                    {{ end }}
                </div>

                <nav id="pagination"></nav>

                </div>
            </div>
            </section>

        </div>

        <script src="https://cdn.jsdelivr.net/npm/flowbite@2.5.2/dist/flowbite.min.js"></script>
        <script>
            function logout() {
                    fetch('/logout', {
                        method: 'DELETE',
                    })
                    .then(response => {
                        if (!response.ok) {
                            throw new Error('Network response was not ok');
                        }
                        window.location.href = "/"
                    })
                    .then(data => {
                        console.log(data); // Process your data here
                    })
                    .catch(error => {
                        console.error('There has been a problem with your fetch operation:', error);
                    });
                }
        </script>
        <script>
            tailwind.config = {
                darkMode: 'media',
                theme: {
                    extend: {
                    colors: {
                        primary: {"50":"#eff6ff","100":"#dbeafe","200":"#bfdbfe","300":"#93c5fd","400":"#60a5fa","500":"#3b82f6","600":"#2563eb","700":"#1d4ed8","800":"#1e40af","900":"#1e3a8a","950":"#172554"}
                    }
                    },
                    fontFamily: {
                    'body': [
                    'Inter', 
                    'ui-sans-serif', 
                    'system-ui', 
                    '-apple-system', 
                    'system-ui', 
                    'Segoe UI', 
                    'Roboto', 
                    'Helvetica Neue', 
                    'Arial', 
                    'Noto Sans', 
                    'sans-serif', 
                    'Apple Color Emoji', 
                    'Segoe UI Emoji', 
                    'Segoe UI Symbol', 
                    'Noto Color Emoji'
                ],
                    'sans': [
                    'Inter', 
                    'ui-sans-serif', 
                    'system-ui', 
                    '-apple-system', 
                    'system-ui', 
                    'Segoe UI', 
                    'Roboto', 
                    'Helvetica Neue', 
                    'Arial', 
                    'Noto Sans', 
                    'sans-serif', 
                    'Apple Color Emoji', 
                    'Segoe UI Emoji', 
                    'Segoe UI Symbol', 
                    'Noto Color Emoji'
                ]
                    }
                }
                }
        </script>
    </body>
{{end}}