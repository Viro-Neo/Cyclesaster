<script lang="ts">
    import { onMount } from 'svelte';
    import L from 'leaflet';
    import 'leaflet/dist/leaflet.css';
    import { Link } from 'svelte-routing';
    import { createEventDispatcher } from 'svelte';

    let map: L.Map;
    let marker: L.Marker;
    let selectedFilter = '';
    const dispatch = createEventDispatcher();
    let franceCoordinates = [46.603354, 1.888334]; // Coordinates for the center of France
    let epitechCoordinates = [48.81549976726981, 2.3629886954816546];

    let filteredList = [
        {
            coordinates: franceCoordinates,
            popupText: 'France center'
        },
        {
            coordinates: epitechCoordinates,
            popupText: 'Epitech Paris'
        }
    ];

    let markers: L.Marker[] = []; // Array to hold markers

    onMount(() => {
        map = L.map('map').setView(franceCoordinates, 6);

        L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
            attribution: '© OpenStreetMap contributors'
        }).addTo(map);

        // Add markers for each item in filteredList on initial mount
        for (const item of filteredList) {
            const marker = L.marker(item.coordinates).addTo(map);
            marker.bindPopup(item.popupText);
            markers.push(marker);
        }
    });

    // Watch for changes in the selectedFilter and update markers accordingly
    $: {
        for (const marker of markers) {
            map.removeLayer(marker); // Remove all existing markers
        }

        markers = []; // Clear the markers array
        let marker: L.Marker;

        switch (selectedFilter) {
            case 'both':
                // Show markers for both coordinates
                for (const item of filteredList) {
                    const marker = L.marker(item.coordinates).addTo(map);
                    marker.bindPopup(item.popupText);
                    markers.push(marker);
                }
                break;
            case 'center_fr':
                // Show marker for France center
                marker = L.marker(franceCoordinates).addTo(map);
                marker.bindPopup('France center');
                markers.push(marker);
                break;
            case 'epitech':
                // Show marker for Epitech Paris
                marker = L.marker(epitechCoordinates).addTo(map);
                marker.bindPopup('Epitech Paris');
                markers.push(marker);
                break;
            default:
                break;
        }
    }
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

    .filter-container {
        width: 30%; /* Adjust width as needed */
        float: right; /* Position the filter container to the right */
        padding: 20px; /* Add padding for spacing */
    }
</style>

<div id="map">
</div>

<div class="button-container">
    <Link to="/chart" class="go-to-chart">Go to Chart</Link>
    <!-- Add more buttons/links as needed -->
</div>

<div class="filter-container">
    <select bind:value={selectedFilter}>
        <option value="both">Both</option>
        <option value="center_fr">France Center</option>
        <option value="epitech">EPITECH</option>
        <!-- Add more filter options as needed -->
    </select>
</div>
