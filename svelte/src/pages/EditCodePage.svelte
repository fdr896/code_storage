<script>
    import { API_URL, languageList, mainPageMode, promise } from '../stores.js';


    export let id;
    export let GetAllCodes;

    let code;
    let language, source, description;

    fetch($API_URL + 'codes/' + id)
    .then(response => {
        console.log(response);
        return response.json();
    })
    .then(response => {
        code = response;

        language = code.language
        source = code.source;
        description = code.description;
    })
    .catch(error => {
        alert(error);
        console.log(error);
    })


    function UpdateCode() {
        fetch($API_URL + 'codes', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                id: id,
                language: language,
                source: source,
                description: description,
                date: code.date
            })
        })
        .then(response => {
            if (!response.ok) {
                console.log(response.statusText);
            }
            return response.json();
        })
        .then(response => {
            console.log(response);

            promise.set(GetAllCodes());
            mainPageMode.set('Codes List');
        }) 
        .catch(error => {
            alert(error);
        });
    }
</script>

<style>
    #description, #language {
        display: inline-block;
        margin-right: 150px;
     }

    .code-form {
        background: #E0E0E0;
        width: 850px;
        height: 600px;
        display: block;
        margin-bottom: 10px;
        border-radius: 5px;
        border: 1px solid #ccc;
        box-shadow: 1px 1px 1px #999;
    }
</style>

<label for="description">Description:</label>
<input
size="30"
id="description"
type="text"
bind:value={description} />

<label for="language">Programming language:</label>
<select id="language" bind:value={language}>
    {#each [...$languageList] as lang}
        <option value={lang[0]}>{lang[1]}</option>
    {/each}
</select>
<hr width="850px" align="left" />

<textarea class="code-form" bind:value={source} />

<button on:click={UpdateCode}>Submit</button>