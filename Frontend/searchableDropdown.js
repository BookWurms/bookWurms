fetch("/api/books/searchableFields")
    .then(response => response.json())

    .then(fields => {
        document.getElementById("dropdownSearchable").innerHTML = ""

        //one field name only
        // document.getElementById("dropdownSearchable").appendChild(document.createElement("Option"))
        // document.getElementById("dropdownSearchable").lastChild.value = fields
        // document.getElementById("dropdownSearchable").lastChild.textContent = fields// 

        // all fields
        fields.forEach(field => {
            document.getElementById("dropdownSearchable").appendChild(document.createElement("Option"))
            document.getElementById("dropdownSearchable").lastChild.value = field
            document.getElementById("dropdownSearchable").lastChild.textContent = field
        })

    })
    .catch(error => console.log(error))
