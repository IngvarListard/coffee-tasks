import sqlite3
import re

select_nodes = """SELECT id, name FROM learning_nodes"""
nodes_to_sort = """SELECT id, name, priority FROM learning_nodes"""
documents_to_sort = """SELECT id, name, priority FROM documents"""

def delete_underscores(db):
    cursor = db.cursor()
    cursor.execute(select_nodes)

    update_cursor = db.cursor()
    sql = """UPDATE learning_nodes
            SET name = ?
            WHERE id = ?"""
    while True:
        result = cursor.fetchone()
        if not result:
            break
        new_name = result[1].replace('_', ' ').title()
        update_cursor.execute(sql, (new_name, result[0]))
        print("NEW NAME", new_name)
    db.commit()


def set_priority(db):
    cursor = db.cursor()
    cursor.execute(nodes_to_sort)

    sql = """UPDATE learning_nodes
            SET priority = ?
            WHERE id = ?"""

    numbers_pattern = re.compile(r"\d*\.\d+|\d+")
    records = cursor.fetchall()
    for result in records:
        if not result:
            break
        try:
            number = next(numbers_pattern.finditer(result[1])).group()
        except StopIteration:
            continue

        print("current priority", result[2])
        priority = float(number)

        if priority % 1 != 0:
            priority = int(number.split('.')[1])
            # priority = int(("%0.2f" % priority).split('.')[1])
        
        cursor.execute(sql, (priority, result[0]))
        print("PRIORITY", priority)
    db.commit()

def set_documents_priority(db):
    cursor = db.cursor()
    cursor.execute(documents_to_sort)

    update_cursor = db.cursor()
    sql = """UPDATE documents
            SET priority = ?
            WHERE id = ?"""

    numbers_pattern = re.compile(r"\d*\.\d+|\d+")
    while True:
        result = cursor.fetchone()
        if not result:
            break
        try:
            number = next(numbers_pattern.finditer(result[1])).group()
        except StopIteration:
            continue

        priority = float(number)
        update_cursor.execute(sql, (priority, result[0]))
        print("PRIORITY", number)
    db.commit()

def main():
    db = sqlite3.connect('courses_bot.db')
    set_priority(db)

if __name__ == "__main__":
    main()