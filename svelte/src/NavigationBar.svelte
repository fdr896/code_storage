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
        font-family: Roboto;
        font-style: normal;
        font-weight: bold;
        font-size: 30px;
        margin-top: 25px;
        text-align: center;
        width: 100%;
        position: fixed;
    }

    .redirect-button {
        position: fixed;
        margin-top: 20px;
    }

    .redirect-button:hover {
        cursor: pointer;
    }

    .add {
        right: 120px;
    }

    .list {
        right: 30px;
    }

    .set {
        left: 30px;
    }
</style>

<nav class="navigation-bar">
    <p class="title">CODE STORAGE</p>
    {#if $userState || $mainPageMode === 'Change Password'}
        <img
        src="img/account_circle.png"
        alt="go to account settings"
        title="wanna change your password?"
        class="redirect-button set"
        on:click={() => { userState.set(false); mainPageMode.set('Change Password') }}
        />

        <img
        src="img/add_code.png"
        alt="add code"
        title="wanna add code?"
        class="redirect-button add"
        on:click={() => { userState.set(true); mainPageMode.set('Add Code')}}
        />

        <img
        src="img/go_to_list.png"
        alt="go to codes list"
        title="wanna look at your codes?"
        class="redirect-button list"
        on:click={() => { userState.set(true); mainPageMode.set('Codes List') }}
        />

    {/if}
</nav>