fetch("/api/books/author")
    .then(response => response.json())
    .then(books => document.getElementById("author").textContent = books[0].Author)
    .catch(error => console.log(error))
