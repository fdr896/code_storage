<script>
    import { API_URL, mainPageMode, languageStyle } from '../stores.js';

    import Highlight from 'svelte-highlight';
    import { github } from 'svelte-highlight/styles';

    import EditCode from './EditCodePage.svelte';
    import Spinner from '../Spinner.svelte';


    export let id;
    async function GetCode(id) {
        const response = await fetch($API_URL + 'codes/' + id);
        const code = await response.json();

        if (response.ok) {
            return code;
        } else {
            throw new Error(code);
        }
    }

    function RedirectToEditCode() {
        mainPageMode.set('Edit Code');
    }
</script>

<svelte:head>{@html github}</svelte:head>

{#await GetCode(id)}
    <Spinner />
{:then code}
    <button
    on:click={RedirectToEditCode}
    >Edit this code</button><br><br>

    <p>{code.description}</p>

    <Highlight
    language={$languageStyle.get(code.language)}
    >{code.source}</Highlight>
{:catch error}
    <script>
        alert(error.message);
    </script>
{/await}