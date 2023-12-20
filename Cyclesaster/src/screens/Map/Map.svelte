<script lang="ts">
    import { onMount } from 'svelte';
    import L from 'leaflet';
    import 'leaflet/dist/leaflet.css';
    import { Link } from 'svelte-routing';

    let map: L.Map;
    let filtersName: string[] = [];
    let filtersValue: string[] = [];
    let selectedFilter1: string = '';
    let selectedFilter2: string = '';
    let franceCoordinates = [46.603354, 1.888334]; // Coordinates for the center of France

    onMount(() => {
        map = L.map('map').setView(franceCoordinates, 6);

        L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
            attribution: '© OpenStreetMap contributors'
        }).addTo(map);
    });
</script>

<div id="map">
</div>

<div class="button-container">
    <Link to="/chart">Go to Chart</Link>
</div>

<div class="filter-container">
    <div class="first-filter">
        <select bind:value={ selectedFilter1 }>
            <option value="">Select a filter</option>
            {#each filtersName as filterName}
                <option value={ filterName }>{ filterName }</option>
            {/each}
        </select>
    </div>

    <div class="second-filter">
        <select bind:value={ selectedFilter2 }>
            <option value="">Select the filter's value</option>
                {#each filtersValue as filterValue}
                    <option value={ filterValue }>{ filterValue }</option>
                {/each}
        </select>
    </div>
</div>

<style>

    @import './Map.css';

</style>
