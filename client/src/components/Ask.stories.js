import Ask from "./Ask";

export default {
    title: 'Ask',
    component: Ask,
};

const Template = (args) => ({
    components: { Ask },
    setup() {
        return { args };
    },
    template: '<Ask v-bind="args"></Ask>',
});

export const Primary = Template.bind({});
Primary.args = {
    primary: true,
    label: 'Button',
};