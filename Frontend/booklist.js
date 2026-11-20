fetch("/api/books")
    .then(dbResponse => dbResponse.json())
    .then(books => {
        document.getElementById("books").textContent = "",

            books.forEach(book => {
                html_book = document.getElementById("books");
                html_book.appendChild(document.createElement("span"));
                html_book.lastChild.textContent = book.Author + " - " + book.Title + " - " + book.ISBN;
                html_book.appendChild(document.createElement("span"));
                html_book.appendChild(document.createElement("br"));
            })
    })
    .catch(error => console.log(error));