fetch("/api/books")
    .then(dbResponse => dbResponse.json())
    // .then(booklist => console.log(booklist[0]))
    .then(booklist => (
        document.getElementById("author").textContent = booklist[0].Author,
        document.getElementById("title").textContent = booklist[0].Title))
    .catch(err => console.log("doomed"));