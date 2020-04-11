<script>
    import NavigationBar from './NavigationBar.svelte';
    import LogInPage from './pages/LogInPage.svelte';
    import MainPage from './MainPage.svelte';

    import { API_URL, userState } from './stores.js';

    fetch($API_URL + 'user', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            login: 'admin'
        })
    });
</script>

<NavigationBar />

{#if !$userState}
    <LogInPage />
{:else}
    <MainPage />
{/if}