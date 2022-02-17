import ResponseGroup from "./ResponseGroup";

export default {
    title: 'ResponseGroup',
    component: ResponseGroup,
};

const Template = (args) => ({
    components: { ResponseGroup },
    setup() {
        return { args };
    },
    template: '<response-group v-bind="args"></response-group>',
});

export const Primary = Template.bind({});
Primary.args = {
    group: {
        value: 1,
        "responses": [
            {
                "id":1,
                "value":1,
                "Student": {
                    "id":2,
                    "first_name":"Jacob",
                    "last_name":"Stevens"
                }
            },
            {
                "id":22,
                "value":1,
                "Student":{
                    "id":1,
                    "first_name":"Sam",
                    "last_name":"Pool"
                }
            }
        ],
        "percent":"40%"
    }
};

export const Secondary = Template.bind({});
Secondary.args = {
    group: {
        value: 2,
        "responses": [
            {
                "id":1,
                "value":2,
                "Student": {
                    "id":2,
                    "first_name":"Jacob",
                    "last_name":"Stevens"
                }
            },
            {
                "id":2,
                "value":2,
                "Student":{
                    "id":1,
                    "first_name":"Sam",
                    "last_name":"Pool"
                }
            },
            {
                "id":3,
                "value":2,
                "Student":{
                    "id":3,
                    "first_name":"Sarah",
                    "last_name":"Suthers"
                }
            }
        ],
        "percent":"40%"
    }
};