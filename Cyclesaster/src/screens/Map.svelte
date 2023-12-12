<script lang="ts">
    import { onMount } from 'svelte';
    import L from 'leaflet';
    import 'leaflet/dist/leaflet.css';
    import { Link } from 'svelte-routing';

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
    /* Ensure the map div takes specific width and height */
    #map {
        height: 70vh; /* Use 70% of the viewport height */
        width: 70%; /* Use 70% of the viewport width */
        float: left; /* Float the map to the left */
    }

    .button-container {
        width: 30%; /* Take 30% of the viewport width */
        float: right; /* Float the container to the right */
        padding: 20px; /* Add padding for spacing */
    }

    /* Style the link/button */
    .go-to-chart {
        display: block; /* Ensure the link/button behaves as a block element */
        background-color: #f0f0f0;
        padding: 10px 20px;
        border-radius: 5px;
        text-decoration: none;
        color: #333;
        margin-bottom: 20px; /* Add margin for spacing between buttons */
    }
</style>

<div id="map">
    <!-- Your map content -->
</div>

<div class="button-container">
    <Link to="/chart" class="go-to-chart">Go to Chart</Link>
    <!-- Add more buttons/links as needed -->
</div>
