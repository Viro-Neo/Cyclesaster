<script lang="ts">
    import { onMount } from 'svelte';
    import Chart from 'chart.js/auto';
    import { Link } from 'svelte-routing';
    import { fetchData, type ApiResponse } from "../../api";

    let filtersApi: ApiResponse;
    let filters: string[] = [];
    let selectedFilter: string = '';
    let selectedFilter2: string = '';
    let selectedFilter3: string = '';

    const tableData = [
        { name: 'A', value: 10 },
        { name: 'B', value: 20 },
        { name: 'C', value: 30 },
        { name: 'D', value: 40 },
        { name: 'E', value: 50 },
    ];

    let chart: Chart;
    const labels = tableData.map((d) => d.name);
    const values = tableData.map((d) => d.value);

    onMount(async () => {
        try {
            filtersApi = await fetchData('/get_filters');
            console.log(filtersApi);
        } catch (error) {
            console.log(error);
        }
        filters = filtersApi.filters;
        const ctx = document.getElementById('chart');
        if (ctx) {
            chart = new Chart(ctx, {
                type: 'bar',
                data: {
                    labels,
                    datasets: [
                        {
                            label: 'Chart Data',
                            data: values,
                            backgroundColor: 'rgba(54, 162, 235, 0.5)', // Adjust color as needed
                            borderColor: 'rgba(54, 162, 235, 1)', // Adjust color as needed
                            borderWidth: 1,
                        }
                    ],
                },
                options: {
                    responsive: true,
                    maintainAspectRatio: false,
                },
            });
        }
    });
</script>

<div id="chart-container">
    <canvas id="chart"></canvas>
    <Link to="/map">Go to Map</Link>
</div>

<div class="filter-container">
    <div class="first-filter">
        <div >
            <select bind:value={ selectedFilter }>
                { #each filters as filter }
                    <option value={ filter }>{ filter }</option>
                { /each }
            </select>
        </div>

        <div >
            <select bind:value={ selectedFilter2 }>
                { #each filters as filter }
                    <option value={ filter }>{ filter }</option>
                { /each }
            </select>
        </div>
    </div>

    <div class="second-filter">
        <select bind:value={ selectedFilter3 }>
            { #each filters as filter }
                <option value={ filter }>{ filter }</option>
            { /each }
        </select>
    </div>
</div>

<style>

    @import './Chart.css';

</style>
