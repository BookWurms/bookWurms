<script>
	import { onMount } from 'svelte';
	/* 
	onMount is sveltes function that says "do this when the page is opened." 
	So an event handler for onPageOpened.
	 */

	let books = $state([]); 
	// we can create variables like this, in our case its a book array
	// State means that the variable is REACTIVE. so the variable redraws the page when it changes.
	let message = $state('');

	/* 
	This one should be self explanitory but async function just means that its asynchronous and
	will wait for the data. in this case if we had a ton of books to load or users, it streams it in the background.
	*/
	async function loadBooks() {
		const response = await fetch('/api/books');
		books = await response.json();
	}

	/* 
	You can bind an event by calling event as a parameter. so the event Upload will now do the following stuff
	in this case its going to register our post request.
	*/
	async function upload(event) {
		event.preventDefault();
		const formData = new FormData(event.currentTarget);
		const response = await fetch('/api/books', { method: 'POST', body: formData });
		message = await response.text();
		await loadBooks();
	}

	//after this is all defined, we just call the get all books API like this.
	onMount(loadBooks);
</script>

<h1>BookWurm</h1>

<h2>GET All Books</h2>
<!-- each just means for each, so for each book, print title and author. else say no books or loading. -->
{#each books as book} 
	<p>{book.Title} by {book.Author}</p>
{:else}
	<p>loading...</p>
{/each}

<h2>Upload a book</h2>
<form onsubmit={upload}>
	<input type="file" name="ebook" accept=".epub" />
	<button type="submit">Upload</button>
</form>

{#if message}
	<p>{message}</p>
{/if}
