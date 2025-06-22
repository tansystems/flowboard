import * as React from 'react';
import { Create, SimpleForm, TextInput, ReferenceInput, SelectInput } from 'react-admin';

export const DealCreate = (props) => (
    <Create {...props} title="Создать сделку">
        <SimpleForm>
            <TextInput source="title" label="Название" />
            <TextInput source="description" label="Описание" />
            <ReferenceInput source="customer_id" reference="customers" label="Клиент">
                <SelectInput optionText="name" />
            </ReferenceInput>
            <ReferenceInput source="status_id" reference="statuses" label="Статус">
                <SelectInput optionText="name" />
            </ReferenceInput>
        </SimpleForm>
    </Create>
); 