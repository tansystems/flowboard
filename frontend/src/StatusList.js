import * as React from 'react';
import { List, Datagrid, TextField, EditButton, DeleteButton, TextInput } from 'react-admin';
import { isAdmin } from './helpers';

const statusFilters = [<TextInput label="Поиск по названию" source="q" alwaysOn key="q" />];

export const StatusList = props => (
    <List {...props} title="Статусы" filters={statusFilters}>
        <Datagrid rowClick="edit">
            <TextField source="id" label="ID" />
            <TextField source="name" label="Название" />
            <TextField source="color" label="Цвет" />
            <EditButton />
            {isAdmin() && <DeleteButton />}
        </Datagrid>
    </List>
); 