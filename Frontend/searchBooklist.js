
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

//assemble GET route with option and text input
function booklist_byOption(field, text) {
    fetch("/api/books/search?field=" + encodeURIComponent(field) +
        "&query_author=" + encodeURIComponent(text))
        .then(response => response.json())
        .then(books => {
            console.log(books)
            html_book = document.getElementById("booksForOneAuthor")
            html_book.textContent = ""
            books.forEach(book => {
                html_book.appendChild(document.createElement("span"))
                html_book.lastChild.textContent = book.Author + " - " + book.Title
                html_book.appendChild(document.createElement("br"))
            })
        })
}


const authorToBeSearchedFor = document.getElementById("query_author")
const fieldToBeSearchedIn = document.getElementById("bookStructureField")

//bookmarked search URL
//get parameters from URL bar
const autofill__query_author = new URLSearchParams(window.location.search).get("query_author")
const autofill__fieldToBeSearchedIn = new URLSearchParams(window.location.search).get("field")

// console.log(autofill__query_author)
// console.log(autofill__fieldToBeSearchedIn)
// console.log(query_author.value)
// console.log(fieldToBeSearchedIn.value)

// use values from URL and show list
// fill authorToBeSearchedFor
if (autofill__query_author) {
    authorToBeSearchedFor.value = autofill__query_author

    //select dropdown option from URL
    if (autofill__fieldToBeSearchedIn) {
        fieldToBeSearchedIn.value = autofill__fieldToBeSearchedIn
    }

    // booklist_byAuthor(authorToBeSearchedFor)
    booklist_byOption(fieldToBeSearchedIn, authorToBeSearchedFor)
}


//get author from input text field "query_author" and on 'enter' show list
//save search URL to history
authorToBeSearchedFor.addEventListener("keydown", function (event) {
    if (event.key === "Enter") {
        console.log(authorToBeSearchedFor.value)
        console.log(fieldToBeSearchedIn.value)

        //change URL in address bar with search query, so we can bookmark it, have it in history 
        // history.pushState(null, "", "/?query_author=" + encodeURIComponent(authorToBeSearchedFor.value))
        // history.pushState(null, "", "/?" + encodeURIComponent(authorToBeSearchedFor.value) + "=" +
        //     encodeURIComponent(authorToBeSearchedFor.value))
        history.pushState(null, "", "/?field=" + encodeURIComponent(fieldToBeSearchedIn.value) +
            "&query_author=" + encodeURIComponent(authorToBeSearchedFor.value))

        // booklist_byAuthor(authorToBeSearchedFor)
        booklist_byOption(fieldToBeSearchedIn.value, authorToBeSearchedFor.value)
    }
})