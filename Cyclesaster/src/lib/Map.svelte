<script lang="ts">
    import { onMount } from 'svelte';
    import L from 'leaflet';
    import 'leaflet/dist/leaflet.css';

    let map: L.Map;
    let marker: L.Marker;

    let franceCoordinates = [46.603354, 1.888334]; // Coordinates for the center of France
    franceCoordinates = [48.81549976726981, 2.3629886954816546];

    onMount(() => {
        map = L.map('map').setView(franceCoordinates, 6); // Center the map on France

        L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
            attribution: '© OpenStreetMap contributors'
        }).addTo(map);

        // Add a marker at the coordinates
        marker = L.marker(franceCoordinates).addTo(map);
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
