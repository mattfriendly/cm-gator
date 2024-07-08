<script>
    import { onMount } from 'svelte';
    import { getUsers } from '$lib/UserService';

    /**
     * @type {string | any[]}
     */
    let users = [];

    onMount(async () => {
        try {
            const data = await getUsers();
            users = data;
        } catch (error) {
            console.error(error);
        }
    });
</script>

<main>
    <h1>List of Users</h1>
    {#if users.length > 0}
        <ul>
            {#each users as user}
                <li>{user.firstname} {user.lastname} - {user.department}</li>
            {/each}
        </ul>
    {:else}
        <p>No users found</p>
    {/if}
</main>