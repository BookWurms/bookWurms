fetch("/api/books")
    .then(dbResponse => dbResponse.json())
    // .then(booklist => console.log(booklist[0]))
    // .then(booklist => 
    // (
    //     document.getElementById("author").textContent = booklist[0].Author,
    //     document.getElementById("title").textContent = booklist[0].Title)
    // )
    .then(books =>
    {
        document.getElementById("books").textContent = "", 

        books.forEach(book => 
        {
            // document.getElementById("books").appendChild(document.createElement("span"));
            // // document.getElementById("books").lastChild.textContent =  book.Author;
            // // document.getElementById("books").appendChild(document.createElement("span"));

            // // document.getElementById("books").appendChild(document.createElement("span"));
            // // document.getElementById("books").lastChild.textContent = " by ";
            // // document.getElementById("books").appendChild(document.createElement("span"));

            // // document.getElementById("books").appendChild(document.createElement("span"));
            // // document.getElementById("books").lastChild.textContent =  book.Title;

            // document.getElementById("books").lastChild.textContent = book.Author + " by " + book.Title;
            // document.getElementById("books").appendChild(document.createElement("span"));
            // document.getElementById("books").appendChild(document.createElement("br"));

            html_book = document.getElementById("books");
            html_book.appendChild(document.createElement("span"));
            html_book.lastChild.textContent = book.Author + " by " + book.Title;
            html_book.appendChild(document.createElement("span"));
            html_book.appendChild(document.createElement("br"));
        })
    })
    // .catch(err => console.log("doomed"));
    .catch(error => console.log(error));