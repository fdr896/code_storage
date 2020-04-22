<script>
    import Spinner from './Spinner.svelte';
    import MainPage from './MainPage.svelte';
    import LogInPage from './pages/LogInPage.svelte';
    import NavigationBar from './NavigationBar.svelte';

    import { API_URL, userState } from './stores.js';

    async function AddUser(user) {
        return await fetch($API_URL + 'user', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                login: user
            })
        });
    }
    let promise = AddUser('admin');
</script>

{#await promise}
    <Spinner />
{:then response}
    <NavigationBar />

    {#if !$userState}
        <LogInPage />
    {:else}
        <MainPage />
    {/if}
{/await}