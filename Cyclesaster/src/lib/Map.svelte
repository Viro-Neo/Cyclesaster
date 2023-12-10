<script lang="ts">
    import { onMount } from 'svelte';
    import L from 'leaflet';
    import 'leaflet/dist/leaflet.css';

    let map: L.Map;

    onMount(() => {
        map = L.map('map').setView([46.603354, 1.888334], 6); // Coordinates for the center of France

        L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
            attribution: '© OpenStreetMap contributors'
        }).addTo(map);
    });
</script>

<style>
    /*!* Set the body and html to take full height *!*/
    /*html, body {*/
    /*    height: 100%;*/
    /*    margin: 0;*/
    /*    padding: 0;*/
    /*}*/

    /* Ensure the map div takes full height */
    #map {
        height: 50vh; /* Use the full viewport height */
        width: 70lh; /* Take full width */
    }
</style>

<link rel="stylesheet" href="https://unpkg.com/leaflet/dist/leaflet.css" />

<div id="map"></div>
