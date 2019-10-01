import { combineReducers } from 'redux'
export default combineReducers({
    num: ChangeNum,
    text: ChangeText,
});

function ChangeNum(state = 0, action: { type: string, num: number }) {
    switch (action.type) {
        case 'add':
            return state += action.num;
        case 'clear':
            return 0;
    }
    return state;
}
function ChangeText(state = '', action: { type: string, text: string }) {
    if (action.type == "append")
        return state += action.text;
    else
        return state
}