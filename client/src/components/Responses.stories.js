import Responses from "./Responses";
import withMock from "storybook-addon-mock";

const mockedCurrentQuestion = { data: { id: 1, text: 'Test Question' } };
const mockedResponses = [{"id":1,"value":1,"Student":{"id":2,"first_name":"Jacob","last_name":"Stevens"}},{"id":2,"value":3,"Student":{"id":3,"first_name":"Sarah","last_name":"Suthers"}},{"id":3,"value":3,"Student":{"id":4,"first_name":"Robert","last_name":"Grainger"}},{"id":21,"value":4,"Student":{"id":1,"first_name":"Sam","last_name":"Pool"}},{"id":22,"value":1,"Student":{"id":1,"first_name":"Sam","last_name":"Pool"}}]

export default {
    title: 'Responses',
    component: Responses,
    decorators: [withMock]
};

const Template = (args) => ({
    components: { Responses },
    setup() {
        return { args };
    },
    template: '<responses v-bind="args"></responses>',
});

export const Primary = Template.bind({});
Primary.args = {
    question: {"id":1,"text":"How ready are you for the math test tomorrow?"},
    responses: [{"id":1,"value":1,"Student":{"id":2,"first_name":"Jacob","last_name":"Stevens"}},{"id":2,"value":3,"Student":{"id":3,"first_name":"Sarah","last_name":"Suthers"}},{"id":3,"value":3,"Student":{"id":4,"first_name":"Robert","last_name":"Grainger"}},{"id":21,"value":4,"Student":{"id":1,"first_name":"Sam","last_name":"Pool"}},{"id":22,"value":1,"Student":{"id":1,"first_name":"Sam","last_name":"Pool"}}]
};
Primary.parameters = {
    mockData: [
        {
            url: '/api/responses/current',
            method: 'GET',
            status: 200,
            response: {"question":{"id":1,"text":"How ready are you for the math test tomorrow?"},"responses":[{"id":1,"value":1,"Student":{"id":2,"first_name":"Jacob","last_name":"Stevens"}},{"id":2,"value":3,"Student":{"id":3,"first_name":"Sarah","last_name":"Suthers"}},{"id":3,"value":3,"Student":{"id":4,"first_name":"Robert","last_name":"Grainger"}},{"id":21,"value":4,"Student":{"id":1,"first_name":"Sam","last_name":"Pool"}},{"id":22,"value":1,"Student":{"id":1,"first_name":"Sam","last_name":"Pool"}}]},
        }
    ]
}
