<script>
    import Spinner from '../Spinner.svelte';
    import { userState, API_URL, mainPageMode } from '../stores.js';
    

    function GoToSignIn(event) {
        if (event.key === 'Enter') {
            SignIn();
        }
    }

    function GoToSignUp(event) {
        if (event.key === 'Enter') {
            SignUp()
        }
    }


    let inputPassword;

    function SignIn() {
        VerifyPassword(inputPassword)
        .then(response => {
            if (response) {
                userState.set(true);
            } else {
                alert('Incorrect password');
                inputPassword = '';
            }
        });
    }

    function SignUp() {
        SetPassword('admin', inputPassword)
        .then(response => {
            if (response.message) {
                alert(response.message);
            } else {
                inputPassword = '';
                promise = HasPassword('admin');

                if ($mainPageMode === 'Change Password') {
                    mainPageMode.set('Codes List');
                    userState.set(true);
                }
            }
        });
    }


    async function HasPassword(user) {
        const response = await fetch($API_URL + 'user/' + user);

        return await response.json();
    }
    let promise = HasPassword('admin');


    async function SetPassword(user, password) {
        const response = await fetch($API_URL + 'user/' + user + '/set', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                password: password
            })
        });

        return await response.json();
    }

    async function VerifyPassword(password) {
        const response = await fetch($API_URL + 'user/admin', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                password: password
            })
        });

        return await response.json();
    }
</script>

<style>
    .password-form {
        margin-top: 15%;
        width: 100%;
        text-align: center;
    }
</style>

<div class="password-form">
    {#await promise}
        <Spinner />
    {:then response}
        {#if response && $mainPageMode !== 'Change Password'}
            <label
            style="margin-top: 0;"
            for="password-field"
            >Password:</label>
            <input
            id="password-field"
            autocomplete="off"
            type="password"
            maxlength="30"
            placeholder="Password"
            bind:value={inputPassword}
            on:keydown={GoToSignIn}
            />

            <button on:click={SignIn}>Sign in</button>
        {:else}
            <label
            for="password-field"
            >Password:</label>
            <input
            id="password-field"
            autocomplete="off"
            type="password"
            maxlength="30"
            placeholder="set password for your account"
            bind:value={inputPassword}
            on:keydown={GoToSignUp}
            />

            <button on:click={SignUp}>Set Password</button>
        {/if}
    {:catch error}
        <script>
            alert(error);
        </script>
    {/await}
</div>