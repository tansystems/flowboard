import * as React from 'react';
import { List, Datagrid, TextField, NumberField, ReferenceField, EditButton, DeleteButton } from 'react-admin';

export const DealList = (props) => (
    <List {...props} title="Сделки">
        <Datagrid rowClick="edit">
            <TextField source="id" label="ID" />
            <TextField source="title" label="Название" />
            <TextField source="description" label="Описание" />
            <ReferenceField source="customer_id" reference="customers" label="Клиент">
                <TextField source="name" />
            </ReferenceField>
            <ReferenceField source="status_id" reference="statuses" label="Статус">
                <TextField source="name" />
            </ReferenceField>
            <EditButton />
            <DeleteButton />
        </Datagrid>
    </List>
); 