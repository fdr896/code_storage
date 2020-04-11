<script>
    import { mainPageMode } from './stores.js';

    import AddCode from './pages/AddCodePage.svelte';
    import CodesList from './pages/CodesListPage.svelte';
    import { ViewPage, AddPage } from './PageMode.js';
</script>

{#if ViewPage($mainPageMode)}
    <CodesList />
{:else if AddPage($mainPageMode)}
    <AddCode />
{/if}
