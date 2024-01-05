const genderMapping: { [key: string]: string } = {
    "-1": "Unknown",
    "1": "Man",
    "2": "Woman"
};

const trafficMapping: { [key: string]: string } = {
    "-1": "Unknown",
    "1": "One way",
    "2": "Two way",
    "3": "One way with divider",
    "4": "Varying directions"
};

const surfaceMapping: { [key: string]: string } = {
    "-1": "Unknown",
    "1": "Normal",
    "2": "Wet",
    "3": "Puddles",
    "4": "Flooded",
    "5": "Snowy",
    "6": "Muddy",
    "7": "Icy",
    "8": "Oily",
    "9": "Other"
};

const infrastructureMapping: { [key: string]: string } = {
    "-1": "Unknown",
    "0": "None",
    "1": "Tunnel - underpass",
    "2": "Bridge",
    "3": "Interchange or connection ramp",
    "4": "Railroad",
    "5": "Crossroads",
    "6": "Pedestrian area",
    "7": "Toll zone",
    "8": "Construction site",
    "9": "Other"
};

const situationMapping: { [key: string]: string } = {
    "-1": "Unknown",
    "0": "None",
    "1": "Road",
    "2": "Emergency lane",
    "3": "Shoulder",
    "4": "Sidewalk",
    "5": "Bike lane",
    "6": "Other special lane",
    "8": "Other",
};

const monthMapping: { [key: string]: string } = {
    "1": "January",
    "2": "February",
    "3": "March",
    "4": "April",
    "5": "May",
    "6": "June",
    "7": "July",
    "8": "August",
    "9": "September",
    "10": "October",
    "11": "November",
    "12": "December"
};

export function mapFilterValue(filterName: string, filterValue: string): string {
    if (filterName === "Gender") {
        return genderMapping[filterValue];
    } else if (filterName === "Trafic") {
        return trafficMapping[filterValue];
    } else if (filterName === "Surface") {
        return surfaceMapping[filterValue];
    } else if (filterName === "Infrastructure") {
        return infrastructureMapping[filterValue];
    } else if (filterName === "Situation") {
        return situationMapping[filterValue];
    } else if (filterName === "Month") {
        return monthMapping[filterValue];
    }
    return filterValue;
}
