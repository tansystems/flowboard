import * as React from 'react';
import { Edit, SimpleForm, TextInput, ReferenceInput, SelectInput } from 'react-admin';

export const DealEdit = (props) => (
    <Edit {...props} title="Редактировать сделку">
        <SimpleForm>
            <TextInput disabled source="id" label="ID" />
            <TextInput source="title" label="Название" />
            <TextInput source="description" label="Описание" />
            <ReferenceInput source="customer_id" reference="customers" label="Клиент">
                <SelectInput optionText="name" />
            </ReferenceInput>
            <ReferenceInput source="status_id" reference="statuses" label="Статус">
                <SelectInput optionText="name" />
            </ReferenceInput>
        </SimpleForm>
    </Edit>
); 