
fetch("/api/books/author")
    .then(response => response.json())

    .then(books => {
        document.getElementById("author").textContent = ""

        books.forEach(book => {
            document.getElementById("author").appendChild(document.createElement("span"))
            document.getElementById("author").lastChild.textContent = book.Author
            document.getElementById("author").appendChild(document.createElement("br"))
        })
    })
    .catch(error => console.log(error))
