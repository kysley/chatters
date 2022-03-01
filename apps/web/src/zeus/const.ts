/* eslint-disable */

export const AllTypesProps: Record<string,any> = {
	Query:{
		chatter:{
			username:{
				type:"String",
				array:false,
				arrayRequired:false,
				required:true
			}
		},
		uses:{
			code:{
				type:"String",
				array:false,
				arrayRequired:false,
				required:true
			}
		}
	}
}

export const ReturnTypes: Record<string,any> = {
	Chatter:{
		id:"ID",
		occurances:"Occurance",
		username:"String"
	},
	Emote:{
		code:"String",
		emoteId:"String",
		id:"ID",
		occurances:"Occurance"
	},
	Occurance:{
		chatter:"Chatter",
		emote:"Emote",
		id:"ID",
		uses:"Int"
	},
	Query:{
		chatter:"Chatter",
		stats:"Stats",
		uses:"Int"
	},
	Stats:{
		chatters:"Int",
		occurances:"Int"
	}
}