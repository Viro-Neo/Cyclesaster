<script lang="ts">
    import { onDestroy, onMount } from 'svelte';
    import Chart from 'chart.js/auto';
    import { Link } from 'svelte-routing';
    import { type ApiResponse, handleFilter, fetchFiltersValues, fetchGraphData } from '../../api';
    import { mapFilterValue } from '../../filterMapping';

    let filtersName: string[] = [];
    let filters: { name: string, value: string, values: string[ ]}[] = [];
    let selectedFilter3: string = '';

    let tableData: Object;
    let tableArray: [string, any][] = [];

    let chart: Chart;
    const labels: string[] = [];
    const values: number[] = [];

    async function handleFilterValue(index: number) {
        try {
            const filterValues = await fetchFiltersValues(filters[index].name);
            filters[index].values = filterValues.filter_values;
        } catch (error) {
            console.log(error);
        }
    }

    export async function handleGraphRequest() {
        try {
            const filterApi = await fetchGraphData(filters, selectedFilter3);
            const canvas = document.getElementById('chart') as HTMLCanvasElement;

            tableData = filterApi.data;
            tableArray = Object.entries(tableData);
            labels.length = 0;
            values.length = 0;
            tableArray.forEach((item) => {
                labels.push(mapFilterValue(selectedFilter3, item[0]));
                values.push(item[1]);
            });
            if (canvas) {
                if (chart) {
                    chart.destroy();
                }
                chart = new Chart(canvas, {
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

    function addFilter() {
        console.log('add filter');
        filters = [...filters, { name: '', value: '', values: [] }];
        console.log(filters);
    }

    onMount(async () => {
        filtersName = await handleFilter();
        console.log(filtersName);
    });

    onDestroy(() => {
        if (chart) {
            chart.destroy();
        }
    });

    $: {
        for (let index = 0; index < filters.length; index++) {
            if (filters[index].name && !filters[index].values.length) {
                handleFilterValue(index);
            }
        }

        if (filters.length >= 1 && selectedFilter3) {
            handleGraphRequest();
        }
    }

</script>

<div id="chart-container">
    <canvas id="chart"></canvas>
    <Link to="/map">Go to Map</Link>
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
