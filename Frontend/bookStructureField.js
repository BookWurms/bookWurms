fetch("/api/books/searchableFields")
    .then(response => response.json())

    .then(fields => {
        document.getElementById("bookStructureField").innerHTML = ""

        //one field name only
        // document.getElementById("dropdownSearchable").appendChild(document.createElement("Option"))
        // document.getElementById("dropdownSearchable").lastChild.value = fields
        // document.getElementById("dropdownSearchable").lastChild.textContent = fields// 

        // all fields
        fields.forEach(field => {
            document.getElementById("bookStructureField").appendChild(document.createElement("Option"))
            document.getElementById("bookStructureField").lastChild.value = field
            document.getElementById("bookStructureField").lastChild.textContent = field
        })

    })
    .catch(error => console.log(error))

