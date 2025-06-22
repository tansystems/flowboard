import * as React from 'react';
import { Edit, SimpleForm, TextInput } from 'react-admin';

export const CustomerEdit = (props) => (
    <Edit {...props} title="Редактировать клиента">
        <SimpleForm>
            <TextInput disabled source="id" label="ID" />
            <TextInput source="name" label="Имя" />
            <TextInput source="email" label="Email" />
            <TextInput source="phone" label="Телефон" />
            <TextInput source="company" label="Компания" />
        </SimpleForm>
    </Edit>
); 