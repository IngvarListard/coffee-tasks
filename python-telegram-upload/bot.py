#!/usr/bin/env python
import logging
import os
import pathlib
import sqlite3
from datetime import datetime
from xml.dom.minidom import DocumentLS

import telegram
from telegram.ext import (CallbackContext, CommandHandler, Filters,
                          MessageHandler, Updater)
from telegram.utils.request import Request
from telegram.ext.dispatcher import run_async

try:
    from secret import REQUEST_KWARGS, TOKEN
except ModuleNotFoundError:
    REQUEST_KWARGS = {}
    TOKEN = ""

logging.basicConfig(
    format='%(asctime)s - %(name)s - %(levelname)s - %(message)s',
    level=logging.INFO)

logger = logging.getLogger(__name__)



class Telegram:
    bot: telegram.Bot = None
    chat_id = None


def start(update: telegram.Update, context: CallbackContext):
    Telegram.chat_id = update.message.chat.id
    handle_dir('./courses', init=True)


def exists(path, type='documents'):
    db = sqlite3.connect('courses_bot.db')
    cursor = db.cursor()
    if type == 'documents':
        cursor.execute("""
            SELECT id FROM documents
            WHERE path=?
        """, (str(path),))
        result = cursor.fetchone()
    elif type == 'node':
        cursor.execute("""
            SELECT id FROM learning_nodes
            WHERE path=?
        """, (str(path),))
        result = cursor.fetchone()

    if result:
        return result[0]
    else:
        return

def get_attachment_from_file(message: telegram.Message):
    required_types = ['audio', 'game', 'animation', 'document', 'photo', 'sticker', 'video',
                        'voice', 'video_note', 'contact', 'location', 'venue', 'invoice',
                        'successful_payment']
    for t in required_types:
        obj = getattr(message, t)
        if obj:
            return obj, t

    print("Неизвестный тип данных", message.effective_attachment(), type(message.effective_attachment()))
    return None, None

    


def handle_document(path, parent_node_id, file_name):
    bot = Telegram.bot
    filename, file_extension = os.path.splitext(path)
    message: telegram.Message = None
    path_ = path
    path = str(pathlib.Path(*pathlib.Path(path).parts[1:]))  # удаление первой папки из пути courses

    with open(path_, 'rb') as f:
        if exists(path):
            print(f'Документ "{file_name} уже загружен"')
            return
        if file_extension == '.mp3':
            message = bot.send_audio(Telegram.chat_id, f, timeout=120)
            create_document(
                file_name.replace('_', ' ').capitalize(),
                file_name,
                message.audio.file_id,
                'audio',
                1,
                parent_node_id,
                path
            )
        elif file_extension in ['.pdf', '.txt']:
            message = bot.send_document(Telegram.chat_id, f, timeout=120)
            create_document(
                file_name.replace('_', ' ').capitalize(),
                file_name,
                message.document.file_id,
                'document',
                2,
                parent_node_id,
                str(path)
            )
    if message:
        print(f'Документ {file_name} отправлен')
        

def handle_dir(path, dirname=None, parent_id=None, init=False):
    node_id = None
    if not init:
        node_id = create_node(dirname, path, parent_id)
    for entry in os.listdir(path):
        new_path = os.path.join(path, entry)
        if os.path.isdir(new_path):
            handle_dir(new_path, entry, node_id)
        else:
            handle_document(new_path, node_id, entry)


def create_node(name: str, path: str, parent_id: int = None):
    db = sqlite3.connect('courses_bot.db')
    cursor = db.cursor()
    path = str(pathlib.Path(*pathlib.Path(path).parts[1:]))  # удаление первой папки из пути courses
    if r := exists(path, 'node'):
        return r

    cursor.execute("""
        SELECT id
        FROM learning_nodes
        WHERE
            name=? AND
            path=? AND
            parent_id=?;
    """, (name.replace('_', ' ').capitalize(), str(path), parent_id))
    r = cursor.fetchall()
    if r:
        if r > 1:
            raise ValueError('multiple items returned, expected one')
        return r[0][0]

    now = datetime.now()
    params = (name, parent_id, name, path, now, now)
    cursor.execute("""
        INSERT INTO learning_nodes
            ("name", "parent_id", "dir_name", "path", "created_at", "updated_at")
        VALUES
            (?,?,?,?,?,?)
    """, params)
    db.commit()
    db.close()
    return cursor.lastrowid


def create_document(name: str, file_name: str, file_id: int, type_: str, priority: float, parent_node_id: int, path: str):
    db = sqlite3.connect('courses_bot.db')
    cursor = db.cursor()
    now = datetime.now()
    params = (name, file_name, file_id, path, type_, priority, parent_node_id, now, now)
    cursor.execute("""
        INSERT INTO documents
            ("name", "file_name", "file_id", "path", "type", "priority", "node_id", "created_at", "updated_at")
        VALUES
            (?,?,?,?,?,?,?,?,?)
    """, params)
    db.commit()
    db.close()
    return cursor.lastrowid


def error(update, context):
    """Log Errors caused by Updates."""
    logger.warning('Update "%s" caused error "%s"', update, context.error)


def hello(u: telegram.Update, c):
    u.message.reply_text('Hi!')

def main():
    updater = Updater(TOKEN, request_kwargs=REQUEST_KWARGS, use_context=True)

    Telegram.bot = updater.bot

    dp = updater.dispatcher
    dp.add_handler(CommandHandler("start", start))
    dp.add_handler(CommandHandler("hello", hello))
    dp.add_error_handler(error)

    updater.start_polling()
    updater.idle()


if __name__ == '__main__':
    main()
