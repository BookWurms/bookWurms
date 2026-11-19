
//call api for booklist by author search
function booklist_byAuthor(authorToBeSearchedFor) {
    //Uniform Resource Identifier
    //Uniform Resource Locator
    //encode 
    //? starts a query parameter
    //= separates parameter name and value
    //& separates multiple parameters
    //%20 represents one space
    //localhost:8443/api/books/author?query_author=Fyodor%20Dostoevsky
    // fetch("/api/books/author?query_author=" + encodeURI("Fyodor Dostoevsky"))
    fetch("/api/books/author?query_author=" + encodeURIComponent(authorToBeSearchedFor.value))
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


const authorToBeSearchedFor = document.getElementById("query_author")
// const fieldToBeSearchedIn = document.getElementById("searchable")

//bookmarked search URL
//get parameters from URL bar
const autofill__query_author = new URLSearchParams(window.location.search).get("query_author")
// const autofill__searchable = new URLSearchParams(window.location.search).get("searchable")
console.log(autofill__query_author)
// console.log(autofill__searchable)
// console.log(autofill_searchable.value)

//fill authorToBeSearchedFor with value from URL and show list
if (autofill__query_author) {
    authorToBeSearchedFor.value = autofill__query_author

    booklist_byAuthor(authorToBeSearchedFor)

}

//get author from input text field "query_author" and on 'enter' show list
//save search URL to history
authorToBeSearchedFor.addEventListener("keydown", function (event) {
    if (event.key === "Enter") {
        console.log(authorToBeSearchedFor.value)

        //change URL in address bar with search query, so we can bookmark it, have it in history 
        // history.pushState(null, "", "/?query_author=" + encodeURIComponent(authorToBeSearchedFor.value))
        history.pushState(null, "", "/?" + encodeURIComponent(searchable.value) + "=" +
            encodeURIComponent(authorToBeSearchedFor.value))

        booklist_byAuthor(authorToBeSearchedFor)
    }
})