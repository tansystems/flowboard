import * as React from 'react';
import { Create, SimpleForm, TextInput } from 'react-admin';

export const StatusCreate = props => (
    <Create {...props} title="Создать статус">
        <SimpleForm>
            <TextInput source="name" label="Название" />
            <TextInput source="color" label="Цвет" />
        </SimpleForm>
    </Create>
); 