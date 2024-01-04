<script lang="ts">
    // Import necessary modules
    import { Link } from 'svelte-routing';

</script>

<div>
    <h1>Welcome!</h1>
    <div>
        <li><Link to="/map">Go to Map</Link></li>
        <li><Link to="/chart">Go to Chart</Link></li>
    </div>
</div>

<style>
    div {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        height: 50vh;
    }

    h1 {
        font-size: 3rem;
        margin-bottom: 2rem;
    }

    li {
        font-size: 2rem;
        margin-bottom: 1rem;
    }
</style>
