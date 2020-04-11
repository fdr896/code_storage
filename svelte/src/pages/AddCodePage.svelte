<script>
    import { mainPageMode, languageList, API_URL } from '../stores.js';


    let source, description, language;


    function CreateCode() {
        fetch($API_URL + 'codes', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                source: source,
                language: language,
                description: description,
            })
        })
        .then(response => {
            console.log(response);
            return response.json();
        })
        .then(response => {
            if (response.message) {
                alert(response.message);
            } else {
                console.log(response);
                mainPageMode.set('Codes List');
            }
        })
        .catch(error => {
            alert(error);
        })
    }
</script>

<style>
    #description, #language {
        display: inline-block;
        margin-right: 70px;
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

<label for="description">Add description:</label>
<input
size="30"
id="description"
type="text"
bind:value={description}
/>

<label for="language">Select programming language:</label>
<select id="language" bind:value={language}>
    {#each [...$languageList] as lang}
        <option value={lang[0]}>{lang[1]}</option>
    {/each}
</select>
<hr width="850px" align="left" />

<textarea
class="code-form"
bind:value={source}
placeholder="Put your code here"
/>

<button on:click={CreateCode}>Create new code</button>