<script lang="ts">
    // Import necessary modules
    import { Link } from 'svelte-routing';

</script>

<div>
    <h1>Welcome!</h1>
    <nav>
        <ul>
            <li><Link to="/map">Go to Map</Link></li>
            <li><Link to="/chart">Go to Chart</Link></li>
        </ul>
    </nav>
</div>
