
// //one specific author
// fetch("/api/books/author")
//     .then(response => response.json())

//     .then(books => {
//         document.getElementById("query_author").textContent = ""

//         books.forEach(book => {
//             document.getElementById("query_uthor").appendChild(document.createElement("span"))
//             document.getElementById("query_author").lastChild.textContent = book.Author + " - " + book.Title
//             document.getElementById("query_author").appendChild(document.createElement("br"))
//         })
//     })
//     .catch(error => console.log(error))


//get author from input text field "query_author"
const authorToBeSearchedFor = document.getElementById("query_author")

// take each input
// authorToBeSearchedFor.addEventListener("input", function () {

//take on enter hole field content
authorToBeSearchedFor.addEventListener("keydown", function (event) {
    if (event.key === "Enter") {
        console.log(authorToBeSearchedFor.value)

        // console.log(authorToBeSearchedFor.value)
        //Uniform Resource Identifier
        //Uniform Resource Locator
        //encode 
        //? starts a query parameter
        //= separates parameter name and value
        //& separates multiple parameters
        //%20 represents one space
        //localhost:8443/api/books/author?query_author=Fyodor%20Dostoevsky
        // fetch("/api/books/author?query_author=" + encodeURI("Fyodor Dostoevsky"))
        fetch("/api/books/author?query_author=" + encodeURI(authorToBeSearchedFor.value))
            .then(response => response.json())

            .then(books => {
                document.getElementById("booksForOneAuthor").textContent = ""

                books.forEach(book => {
                    document.getElementById("booksForOneAuthor").appendChild(document.createElement("span"))
                    document.getElementById("booksForOneAuthor").lastChild.textContent = book.Author + " - " + book.Title
                    document.getElementById("booksForOneAuthor").appendChild(document.createElement("br"))
                })
            })
            .catch(error => console.log(error))
    }
})