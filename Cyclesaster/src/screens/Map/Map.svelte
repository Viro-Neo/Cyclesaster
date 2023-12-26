<script lang="ts">
    import { onMount } from 'svelte';
    import L from 'leaflet';
    import 'leaflet/dist/leaflet.css';
    import { Link } from 'svelte-routing';
    import { type ApiResponse, handleFilter, fetchFiltersValues, fetchMapData, fetchAccidentById } from '../../api';
    import { mapFilterValue } from '../../filterMapping';

    let map: L.Map;
    let filtersApi2: ApiResponse;
    let filtersName: string[] = [];
    let filtersValue: string[] = [];
    let yearApi: ApiResponse;
    let yearsFilter: string[] = [];
    let selectedFilter: string = '';
    let selectedFilter2: string = '';
    let selectedYear: string = '';
    let franceCoordinates = [46.603354, 1.888334]; // Coordinates for the center of France

    async function handleFilterValue() {
        try {
            filtersApi2 = await fetchFiltersValues(selectedFilter);
            filtersValue = filtersApi2.filter_values;
        } catch (error) {
            console.log(error);
        }
    }

    interface AccidentData {
        "Id": number;
        "Day": string;
        "Month": string;
        "Year": string;
        "Birth_year": string,
        "Departement": string;
        "Gender": string;
        "Surface": string;
        "Infrastructure": string;
        "Trafic": string;
        "Situation": string;
        "Latitude": string;
        "Longitude": string;

        [key: string]: any;
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

    async function handleMapRequest(selectedFilter: string, selectedFilter2: string) {
        console.log(selectedFilter, selectedFilter2);
        try {
            console.log("fetching map data");
            console.log("selectedYear", selectedYear)
            const filterApi = await fetchMapData(selectedFilter, selectedFilter2, selectedYear);
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

    onMount(async () => {
        filtersName = await handleFilter(filtersName);
        yearApi = await fetchFiltersValues('Year');
        yearsFilter = yearApi.filter_values;
        map = L.map('map').setView(franceCoordinates, 6);

        L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
            attribution: '© OpenStreetMap contributors'
        }).addTo(map);
    });

    $: {
        if (selectedFilter) {
            handleFilterValue();
        }

        if (selectedFilter && selectedFilter2 && selectedYear) {
            handleMapRequest(selectedFilter, selectedFilter2);
        }
    }
</script>

<div id="map">
</div>

<div class="button-container">
    <Link to="/chart">Go to Chart</Link>
</div>

<div class="filter-container">
    <div class="first-filter">
        <select bind:value={ selectedFilter }>
            <option value="">Select a filter</option>
            {#each filtersName as filterName}
                <option value={ filterName }>{ filterName }</option>
            {/each}
        </select>
    </div>

    <div class="second-filter">
        <select bind:value={ selectedFilter2 }>
            <option value="">Select the filter's value</option>
                {#each filtersValue.sort((one, two) => (one > two ? -1 : 1)) as filterValue}
                    <option value={ filterValue }>{ mapFilterValue(selectedFilter, filterValue) }</option>
                {/each}
        </select>
    </div>

    <div class="year-filter">
        <select bind:value= { selectedYear }>
            <option value="">Select a year</option>
            {#each yearsFilter.sort((one, two) => (one > two ? -1 : 1)) as filterValue}
                <option value={ filterValue }>{ filterValue }</option>
            {/each}
        </select>
    </div>
</div>

<style>

    @import './Map.css';

</style>
