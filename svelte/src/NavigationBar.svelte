<script>
    import { mainPageMode, userState } from './stores.js';
</script>

<style>
    .navigation-bar {
        z-index: 7474;
        left: 0;
        right: 0;
        height: 90px;
        top: 0;
        position: fixed;
        background-color: #C0C0C0;
        margin-bottom: 20px;
    }

    .title {
        font-family: Arial, Helvetica, sans-serif;
        font-size: 24px;
        margin-top: 30px;
        text-align: center;
        width: 100%;
        position: fixed;
    }

    .redirect-button {
        position: fixed;
        margin-top: 30px;
    }

    .add {
        right: 10px;
    }

    .list {
        right: 130px;
    }

    .set {
        left: 10px;
    }
</style>

<nav class="navigation-bar">
    <p class="title">CODE STORAGE</p>
    {#if $userState || $mainPageMode === 'Change Password'}
        <button
        class="redirect-button set"
        on:click={() => { userState.set(false); mainPageMode.set('Change Password') }}
        >Change Password</button>

        <button
        class="redirect-button add"
        on:click={() => { userState.set(true); mainPageMode.set('Add Code')}}
        >Add code</button>

        <button
        class="redirect-button list"
        on:click={() => { userState.set(true); mainPageMode.set('Codes List') }}
        >Codes list</button>
    {/if}
</nav>