import * as React from 'react';
import { List, Datagrid, TextField, EditButton, DeleteButton, TextInput } from 'react-admin';
import { isAdmin } from './helpers';

const tagFilters = [<TextInput label="Поиск по тегу" source="q" alwaysOn key="q" />];

export const TagList = props => (
    <List {...props} title="Теги" filters={tagFilters}>
        <Datagrid rowClick="edit">
            <TextField source="id" label="ID" />
            <TextField source="name" label="Название" />
            <EditButton />
            {isAdmin() && <DeleteButton />}
        </Datagrid>
    </List>
); 