<script lang="ts">
    import { onDestroy, onMount } from 'svelte';
    import L from 'leaflet';
    import 'leaflet/dist/leaflet.css';
    import { Link } from 'svelte-routing';
    import { type ApiResponse, handleFilter, fetchFiltersValues, fetchMapData, fetchAccidentById } from '../../api';
    import { mapFilterValue } from '../../filterMapping';

    let map: L.Map;
    let filtersName: string[] = [];
    let filters: { name: string, value: string, values: string[] }[] = [];
    let franceCoordinates = [46.603354, 1.888334]; // Coordinates for the center of France
    let shouldHandleMapRequest = false;

    async function handleFilterValue(index: number) {
        try {
            const filterValues = await fetchFiltersValues(filters[index].name);
            filters[index].values = filterValues.filter_values;
        } catch (error) {
            console.log(error);
        }
    }

    function displayData(data, popup: HTMLElement) {
        popup.innerHTML = '';

        const list = document.createElement('ul');

        function processProperty(key, value, parentElement) {
            const listItem = document.createElement('li');

            listItem.style.padding = '0';
            listItem.style.margin = '0';

            if (typeof value === 'object' && value !== null) {
                const sublist = document.createElement('ul');
                sublist.style.padding = '0';
                listItem.innerHTML = `<strong>${key}:</strong>`;
                listItem.appendChild(sublist);

                for (const subKey in value) {
                    if (value.hasOwnProperty(subKey)) {
                        processProperty(subKey, value[subKey], sublist);
                    }
                }
            } else {
                if (value !== "") {
                    listItem.innerHTML = `<strong>${key}:</strong> ${mapFilterValue(key, value)}`;
                } else {
                    listItem.innerHTML = `<strong>${key}:</strong> ${"Unknown"}`;
                }                
            }

            parentElement.appendChild(listItem);
        }

        for (const key in data) {
            if (data.hasOwnProperty(key)) {
                processProperty(key, data[key], list);
            }
        }

        popup.appendChild(list);
    }


    async function fetchDataById(id: number, popup: HTMLElement) {
        try {
            const accidentApi = await fetchAccidentById(id);
            const data = accidentApi;
            displayData(data, popup);
        } catch (error) {
            console.log(error);
        }
    }

    async function handleMapRequest() {
        try {
            const filterApi = await fetchMapData(filters);
            const data = filterApi.data;
            
            if (Array.isArray(data)) {
                data.forEach(item => {
                console.log("adding marker for item", item);
                console.log("latitude", item.Latitude);
                console.log("longitude", item.Longitude);

                const popup = document.createElement('div');
                popup.style.width = '150px';
                popup.style.height = '315px';
                popup.style.padding = '2px';

                const marker = L.marker([parseFloat(item.Latitude), parseFloat(item.Longitude)])
                    .addTo(map)
                    .bindPopup(popup)
                    .on('click', () => {
                        fetchDataById(item.Id, popup);
                    });
                });
            }
        } catch (error) {
            console.log(error);
        }
    }

    function addFilter() {
        filters = [...filters, { name: '', value: '', values: [] }];
    }

    onMount(async () => {
        filtersName = await handleFilter();
        filters = [{ name: 'Year', value: '', values: [] }];
        map = L.map('map').setView(franceCoordinates, 6);

        L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
            attribution: '© OpenStreetMap contributors'
        }).addTo(map);
    });

    onDestroy(() => {
        if (map) {
            map.remove();
        }
    });

    $: {
        for (let index = 0; index < filters.length; index++) {
            if (filters[index].name && !filters[index].values.length) {
                handleFilterValue(index);
            }
        }

        if (filters.length >= 2) {
            shouldHandleMapRequest = true;
            filters.forEach(element => {
                if (element.value == "") {
                    shouldHandleMapRequest = false;
                }
            });
            if (shouldHandleMapRequest)
                handleMapRequest();
        }
    }

</script>

<div id="map">
</div>

<div class="button-container">
    <Link to="/chart">Go to Chart</Link>
</div>

<div class="filter-container">
    {#each filters as filter, index}
        <div class="filter">
            <div>
                <select bind:value={filter.name}>
                    <option value="">Select a filter</option>
                    {#each filtersName as filterName}
                        <option value={filterName}>{filterName}</option>
                    {/each}
                </select>
            </div>

            <div>
                <select bind:value={filter.value}>
                    <option value="">Select the filter's value</option>
                    {#each (filter.values || []).sort((one, two) => (one > two ? -1 : 1)) as filterValue}
                        <option value={filterValue}>{mapFilterValue(filter.name, filterValue)}</option>
                    {/each}
                </select>
            </div>
            {#if index > 0}
            <button on:click={() => { filters = filters.filter((_, i) => i !== index); }}>Remove Filter</button>
            {/if}
        </div>
    {/each}

    <button on:click={addFilter}>{`Add Filter`}</button>
</div>

<style>

    #map {
        height: 70vh; /* Use 70% of the viewport height */
        width: 100%; /* Use 70% of the viewport width */
        float: left; /* Float the map to the left */
    }

    .button-container {
        width: 30%; /* Take 30% of the viewport width */
        float: right; /* Float the container to the right */
        padding: 20px; /* Add padding for spacing */
    }

    .filter-container {
        width: 30%; /* Take 30% of the viewport width */
        float: right; /* Float the container to the right */
        padding: 20px; /* Add padding for spacing */
        border: 1px solid black;
        border-radius: 10px;
    }

    .filter {
        display: flex;
        flex-direction: row;
        justify-content: space-between;
        align-items: center;
        margin-bottom: 10px;
    }


</style>
