<script>
    import { API_URL, mainPageMode, promise } from '../stores.js';

    import ShowCode from './ShowCodePage.svelte';
    import EditCode from './EditCodePage.svelte';
    import Spinner from '../Spinner.svelte';


    async function GetAllCodes() {
        const response = await fetch($API_URL + 'codes');
        const codes = await response.json();

        if (response.ok) {
            codes.sort((i, j) => {
                return i.id > j.id;
            });

            console.log('***CODES LIST***');
            codes.forEach(el => {
                console.log(el);
            })

            return codes;
        } else {
            throw new Error(codes);
        }
    }
    promise.set(GetAllCodes());


    function DeleteCode(id) {
        fetch($API_URL + 'codes/' + id, {
            method: 'DELETE'
        })
        .then(response => {
            if (!response.ok) {
                console.log(response.statusText);
            } else {
                promise.set(GetAllCodes());
            }
        })
        .catch(error => {
            alert(error.message);
        });
    }


    let codeID;
    function RedirectToEditCode(id) {
        codeID = id;
        mainPageMode.set('Show Code');
    }


    function handleText(s) {
        if (s.length > 33) {
            s = s.slice(0, 33) + '...';
        }

        return s;
    }
</script>

{#if $mainPageMode === 'Codes List'}
    {#await $promise}
        <Spinner />
    {:then codes}
        {#if codes.length}
            <style>
                .codes-list {
                    display: grid;
                    width: 100%;
                    grid-row-gap: 20px;
                    grid-column-gap: 10px;
                    justify-items: center;
                    justify-content: center;
                }

                span {
                    background-color: #B0B0B0;
                    border-radius: 5px;
                    line-height: 65px;
                    padding-left: 10px;
                    box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.25);
                    font-family: Arial, Helvetica, sans-serif;
                }

                span:hover {
                    box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.75);
                }

                .grid-item {
                    width: 400px;
                    /* background: linear-gradient(to right, #e0098e, #0720b3) !important; */
                }

                .delete-button {
                    margin-right: 140px;
                    height: 30px;
                    margin-top: 20px;
                }
            </style>

            <div
            class="codes-list"
            style="grid-template-rows: repeat({codes.length}, 70px); grid-template-columns: 400px 30px;"
            >
                {#each codes as code}
                    <span
                    class="grid-item"
                    on:click={() => RedirectToEditCode(code.id)}
                    >{handleText(code.description)}</span>

                    <button
                    class="delete-button"
                    on:click={() => DeleteCode(code.id)}
                    >Delete</button>
                {/each}
            </div>
        {:else}
            <style>
                .empty-list {
                    text-align: center;
                    font-size: 30px;
                }
            </style>

            <p class="empty-list">Codes list is still empty? How dare you?!😈</p>
        {/if}
    {:catch error}
        <script>
            alert(error.message);
        </script>
    {/await}
{:else if $mainPageMode === 'Show Code'}
    <ShowCode id={codeID} />
{:else if $mainPageMode === 'Edit Code'}
    <EditCode id={codeID} GetAllCodes={GetAllCodes} />
{/if}