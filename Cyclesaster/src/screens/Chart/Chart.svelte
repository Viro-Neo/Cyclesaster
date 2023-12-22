<script lang="ts">
    import { onMount } from 'svelte';
    import Chart from 'chart.js/auto';
    import { Link } from 'svelte-routing';
    import { type ApiResponse, handleFilter, fetchFiltersValues, fetchGraphData } from '../../api';
    import { mapFilterValue } from '../../filterMapping';

    let filtersApi2: ApiResponse;
    let filtersName: string[] = [];
    let filtersValue: string[] = [];
    let selectedFilter: string = '';
    let selectedFilter2: string = '';
    let selectedFilter3: string = '';

    let tableData: Object;
    let tableArray: [string, any][] = [];

    let chart: Chart;
    const labels: string[] = [];
    const values: number[] = [];

    async function handleFilterValue() {
        try {
            filtersApi2 = await fetchFiltersValues(selectedFilter);
            filtersValue = filtersApi2.filter_values;
        } catch (error) {
            console.log(error);
        }
    }

    export async function handleGraphRequest(selectedFilter: string, selectedFilter2: string, selectedFilter3: string) {
        try {
            const filterApi = await fetchGraphData(selectedFilter, selectedFilter2, selectedFilter3);
            const ctx = document.getElementById('chart');

            tableData = filterApi.data;
            tableArray = Object.entries(tableData);
            tableArray.forEach((item) => {
                labels.push(mapFilterValue(selectedFilter3, item[0]));
                values.push(item[1]);
            });
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
        } catch (error) {
            console.log(error);
        }
    }

    onMount(async () => {
        filtersName = await handleFilter(filtersName);
    });

    $: {
        if (selectedFilter) {
            handleFilterValue();
        }

        if (selectedFilter && selectedFilter2 && selectedFilter3) {
            handleGraphRequest(selectedFilter, selectedFilter2, selectedFilter3);
        }
    }

</script>

<div id="chart-container">
    <canvas id="chart"></canvas>
    <Link to="/map">Go to Map</Link>
</div>

<div class="filter-container">
    <div class="first-filter">
        <div >
            <select bind:value={ selectedFilter }>
                <option value="">Select a filter</option>
                { #each filtersName as filter }
                    <option value={ filter }>{ filter }</option>
                { /each }
            </select>
        </div>

        <div >
            <select bind:value={ selectedFilter2 }>
                <option value="">Select the filter's value</option>
                { #each filtersValue.sort((one, two) => (one > two ? -1 : 1)) as filter }
                    <option value={ filter }>{ mapFilterValue(selectedFilter, filter) }</option>
                { /each }
            </select>
        </div>
    </div>

    <div class="second-filter">
        <select bind:value={ selectedFilter3 }>
            <option value="">Select a filter</option>
            { #each filtersName as filter }
                <option value={ filter }>{ filter }</option>
            { /each }
        </select>
    </div>
</div>

<style>

    @import './Chart.css';

</style>
