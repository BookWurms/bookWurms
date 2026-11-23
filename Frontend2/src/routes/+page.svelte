<script>
	import { onMount } from "svelte";
	/* 
	onMount is sveltes function that says "do this when the page is opened." 
	So an event handler for onPageOpened.
	 */

	let books = $state([]);
	// we can create variables like this, in our case its a book array
	// State means that the variable is REACTIVE. so the variable redraws the page when it changes.
	let message = $state("");

	/* 
	This one should be self explanitory but async function just means that its asynchronous and
	will wait for the data. in this case if we had a ton of books to load or users, it streams it in the background.
	*/
	async function loadBooks() {
		const response = await fetch("/api/books");
		books = await response.json();
	}

	/* 
	You can bind an event by calling event as a parameter. so the event Upload will now do the following stuff
	in this case its going to register our post request.
	*/
	async function upload(event) {
		event.preventDefault();
		const formData = new FormData(event.currentTarget);
		const response = await fetch("/api/books", {
			method: "POST",
			body: formData,
		});
		message = await response.text();
		await loadBooks();
	}

	// Dropdown options of fields of structure books, that are tagged as 'searchable'
	let searchableFields = $state([]);
	//chosen option, the field we search in for our text
	let optionField = $state("");
	async function dropdownSearchableFields() {
		const response = await fetch("/api/books/searchableFields");
		searchableFields = await response.json();
	}

	//text input field for text to be searched for
	let search_text = $state("");
	//booklist as result of our search
	let booksFound = $state([]);
	async function searchForBooks() {
		console.log(search_text);
		const response = await fetch(
			"/api/books/search?field=" +
				encodeURIComponent(optionField) +
				"&search_text=" +
				encodeURIComponent(search_text),
		);

		console.log(response.status);
		booksFound = await response.json();
		console.log(" ", booksFound);

		//chabge URL
		history.pushState(
			null,
			"",
			"/?field=" +
				encodeURIComponent(optionField) +
				"&search_text=" +
				encodeURIComponent(search_text),
		);
	}

	//after this is all defined, we just call the get all books API like this.
	// onMount(loadBooks);
	onMount(() => {
		loadBooks();
		dropdownSearchableFields();

		//bookmarked search URL
		//get parameters from URL bar
		const autofill__search_text = new URLSearchParams(
			window.location.search,
		).get("search_text");
		const autofill__optionField = new URLSearchParams(
			window.location.search,
		).get("field");

		// use values from URL and show list
		// fill authorToBeSearchedFor
		if (autofill__search_text) {
			search_text = autofill__search_text;

			//select dropdown option from URL
			if (autofill__optionField) {
				optionField = autofill__optionField;
			}
		}
		// searchForBooks();
	});
</script>

<h1>BookWurm</h1>

<h2>GET All Books</h2>
<!-- each just means for each, so for each book, print title and author. else say no books or loading. -->
{#each books as book}
	<p>{book.Title} by {book.Author}</p>
{:else}
	<p>loading...</p>
{/each}

<h2>GET Books by Search</h2>
<!-- dropdown for searchable fields/columns -->
<!-- GORM field tag 'searchable' -->
<select id="bookStructureField" bind:value={optionField}>
	<option value="">Select a field</option>
	{#each searchableFields as field}
		<option value={field}>{field}</option>
	{/each}
</select>
<!-- text input field -->
<input
	type="text"
	bind:value={search_text}
	placeholder="Enter name"
	onkeydown={(event) => {
		if (event.key === "Enter") {
			searchForBooks();
		}
	}}
/>

<!-- call async function to get database results via /api/books/search -->
<button onclick={searchForBooks}>Submit</button>
{#each booksFound as book}
	<p>{book.Title} by {book.Author}</p>
{/each}

<h2>Upload a book</h2>
<form onsubmit={upload}>
	<input type="file" name="ebook" accept=".epub" />
	<button type="submit">Upload</button>
</form>

{#if message}
	<p>{message}</p>
{/if}
