import * as React from 'react';
import { Edit, SimpleForm, TextInput } from 'react-admin';

export const StatusEdit = props => (
    <Edit {...props} title="Редактировать статус">
        <SimpleForm>
            <TextInput disabled source="id" label="ID" />
            <TextInput source="name" label="Название" />
            <TextInput source="color" label="Цвет" />
        </SimpleForm>
    </Edit>
); 