import csv
import argparse
import psycopg2
from psycopg2 import sql

DB_NAME = "cyclesaster_data"
DB_USER = "postgres"
DB_PASSWORD = "postgres"
DB_HOST = "localhost"
DB_PORT = 5432

CSV_FILE_PATH = "data.csv"

COLUMN_MAPPING_CHARACTERISTICS = {
    "Accident_id": ('accident_id', 'bigint'),
    "jour": ('day', 'int'),
    "mois": ('month', 'int'),
    "an": ('short_year', 'int'),
    "hrmn": ('time', 'varchar'),
    "lum": ('lighting', 'int'),
    "dep": ('department', 'varchar'),
    "com": ('municipality', 'varchar'),
    "agg": ('agglomeration', 'int'),
    "int": ('intersection', 'int'),
    "atm": ('weather', 'int'),
    "col": ('collision', 'int'),
    "adr": ('address', 'varchar'),
    "lat": ('latitude', 'float'),
    "long": ('longitude', 'float'),
    "gps": ('gps', 'int')
}

COLUMN_MAPPING_AREA = {
    "Num_Acc": ('accident_id', 'bigint'),
    "catr": ('road_category', 'int'),
    "voie": ('road_number', 'varchar'),
    "v1": ('road_index_1', 'varchar'),
    "v2": ('road_index_2', 'varchar'),
    "circ": ('traffic', 'int'),
    "nbv": ('number_of_lanes', 'int'),
    "vosp": ('bike_lane', 'int'),
    "prof": ('road_profile', 'int'),
    "pr": ('distance', 'varchar'),
    "pr1": ('metered_distance', 'varchar'),
    "plan": ('mapping', 'int'),
    "lartpc": ('road_width', 'float'),
    "larrout": ('roadway_width', 'float'),
    "surf": ('surface', 'int'),
    "infra": ('infrastructure', 'int'),
    "situ": ('situation', 'int'),
    "vma": ('max_speed', 'int'),
    "env1": ('environment', 'int'),
}

COLUMN_MAPPING_VEHICLES = {
    "Num_Acc": ('accident_id', 'bigint'),
    "id_vehicule": ('vehicle_id', 'bigint'),
    "num_veh": ('vehicle_number', 'varchar'),
    "senc": ('direction', 'int'),
    "catv": ('vehicle_category', 'int'),
    "obs": ('mobile_obstacle', 'int'),
    "obsm": ('fixed_obstacle', 'int'),
    "choc": ('impact', 'int'),
    "manv" : ('maneuver', 'int'),
    "motor": ('motorization', 'int'),
    "occutc": ('number_of_occupants', 'int'),
}

COLUMN_MAPPING_USERS = {
    "Num_Acc": ('accident_id', 'bigint'),
    "id_usager": ('user_id', 'bigint'),
    "id_vehicule": ('vehicle_id', 'bigint'),
    "num_veh": ('vehicle_number', 'varchar'),
    "place": ('place', 'int'),
    "catu": ('user_category', 'int'),
    "grav": ('severity', 'int'),
    "sexe": ('gender', 'int'),
    "an_nais": ('birth_year', 'int'),
    "trajet": ('routing', 'int'),
    "secu1": ('safety_equipment_1', 'int'),
    "secu2": ('safety_equipment_2', 'int'),
    "secu3": ('safety_equipment_3', 'int'),
    "locp": ('location', 'int'),
    "actp": ('action', 'varchar'),
    "etatp": ('state', 'int'),
}

def import_csv_data(filename, table_name):
    try:
        conn = psycopg2.connect(
            dbname=DB_NAME,
            user=DB_USER,
            password=DB_PASSWORD,
            host=DB_HOST,
            port=DB_PORT
        )
        cursor = conn.cursor()

        year = filename.split('/')[-1].split('.')[0]
        year = year.split('-')[-1]
        print(year)

        with open(filename, 'r') as file:
            print("Importing data from {} to {} table".format(filename, table_name))
            reader = csv.reader(file)
            print("Reading CSV file")
            header = next(reader)
            print("Creating table")

            column_mapping = {}

            if table_name == "characteristics":
                column_mapping = COLUMN_MAPPING_CHARACTERISTICS
            elif table_name == "area":
                column_mapping = COLUMN_MAPPING_AREA
            elif table_name == "vehicles":
                column_mapping = COLUMN_MAPPING_VEHICLES
            elif table_name == "users":
                column_mapping = COLUMN_MAPPING_USERS

            print(column_mapping)
            print(column_mapping.get("surf", "surf"))

            create_table_query = sql.SQL("CREATE TABLE IF NOT EXISTS {} ({}, year INTEGER)").format(
                sql.Identifier(table_name),
                sql.SQL(', ').join(
                    sql.SQL("{} {}").format(
                        sql.Identifier(mapping[0]),  # column name
                        sql.SQL(mapping[1])  # data type
                    )
                    for column, mapping in column_mapping.items()
                )
            )
            print(create_table_query.as_string(conn))
            cursor.execute(create_table_query)
            print("Table created")

            insert_query = sql.SQL("INSERT INTO {} ({}, year) VALUES ({}, %s)").format(
                sql.Identifier(table_name),
                sql.SQL(', ').join(sql.Identifier(column_mapping.get(column, column)[0]) for column in header),
                sql.SQL(', ').join(sql.Placeholder() for _ in header)
            )

            i = 0
            for row in reader:
                print("Inserting row {}".format(i))
                print("columns are {}".format(header))
                row_with_dots = [value.replace(',', '.') for value in row]
                row_with_none = [None if value == '' else value for value in row_with_dots]
                print(insert_query.as_string(conn))
                cursor.execute(insert_query, row_with_none + [year])
                i += 1
        
        conn.commit()
    
    except psycopg2.Error as e:
        print("Error: Could not make connection to the Postgres database")
        print(e)
    
    finally:

        if conn:
            cursor.close()
            conn.close()

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description='Export CSV data to Postgres database')
    parser.add_argument('filename', type=str, help='CSV file path')
    parser.add_argument('table_name', type=str, help='Table name')

    args = parser.parse_args()

    import_csv_data(args.filename, args.table_name)
