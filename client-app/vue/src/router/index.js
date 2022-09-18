import Vue from 'vue'
import VueRouter from 'vue-router'

import Login from '../views/Login'
import Logout from '../views/Logout'
import UnregisteredPage from '../views/UnregisteredPage'
import RegisteredPage from '../views/RegisteredPage'
import AdministratorPage from '../views/AdministratorPage'
import SupportPage from '../views/SupportPage'

import Register from '../components/Register'
import SearchVideos from '../components/SearchVideos'
import VideoView from '../components/VideoView'
import UploadVideo from '../components/UploadVideo'
import ReportedVideos from '../components/ReportedVideos'
import ReportedComments from '../components/ReportedComments'
import Profile from '../components/Profile'
import Messages from '../components/Messages'
import SupportMessages from '../components/SupportMessages'

Vue.use(VueRouter)

const Role = {
	Administrator: 'Administrator',
	RegisteredUser: 'RegisteredUser',
	SupportUser: 'SupportUser'
}

const routes = [
	{
		path: "/",
		name: UnregisteredPage,
		component: UnregisteredPage,
		children: [
			{
				path: "",
				name: "SearchVideosUnregisteredUser",
				component: SearchVideos
			},
			{
				path: "Register",
				name: "Register",
				component: Register
			},
			{
				path: "VideoView",
				name: "VideoViewUnregisteredUser",
				component: VideoView
			},
		]
	},
	{
		path: "/RegisteredPage",
		name: RegisteredPage,
		component: RegisteredPage,
		children: [
			{
				path: "",
				name: "SearchVideosRegisteredUser",
				component: SearchVideos,
				meta: {
					roles: [Role.RegisteredUser]
				},
			},
			{
				path: "VideoView",
				name: "VideoViewRegisteredUser",
				component: VideoView,
				meta: {
					roles: [Role.RegisteredUser]
				},
			},
			{
				path: "UploadVideo",
				name: "UploadVideo",
				component: UploadVideo,
				meta: {
					roles: [Role.RegisteredUser]
				},
			},
			{
				path: "Profile",
				name: "ProfileRegisteredUser",
				component: Profile,
				meta: {
					roles: [Role.RegisteredUser]
				},
			},
			{
				path: "Messages",
				name: "RegisteredUserMessages",
				component: Messages,
				meta: {
					roles: [Role.RegisteredUser]
				},
			},
		],
		meta: {
			roles: [Role.RegisteredUser]
		},
	},
	{
		path: "/AdministratorPage",
		name: AdministratorPage,
		component: AdministratorPage,
		children: [
			{
				path: "",
				name: "SearchVideosAdministrator",
				component: SearchVideos,
				meta: {
					roles: [Role.Administrator]
				},
			},
			{
				path: "VideoView",
				name: "VideoViewAdministrator",
				component: VideoView,
				meta: {
					roles: [Role.Administrator]
				},
			},
			{
				path: "ReportedVideos",
				name: "ReportedVideos",
				component: ReportedVideos,
				meta: {
					roles: [Role.Administrator]
				},
			},
			{
				path: "ReportedComments",
				name: "ReportedComments",
				component: ReportedComments,
				meta: {
					roles: [Role.Administrator]
				},
			},
			{
				path: "Profile",
				name: "ProfileAdministrator",
				component: Profile,
				meta: {
					roles: [Role.Administrator]
				},
			},
		],
		meta: {
			roles: [Role.Administrator]
		},
	},
	{
		path: "/SupportPage",
		name: SupportPage,
		component: SupportPage,
		children: [
			{
				path: "SupportMessages",
				name: "SupportUserSupportMessages",
				component: SupportMessages,
				children: [
					{
						path: "Messages",
						name: "SupportUserMessages",
						component: Messages,
						meta: {
							roles: [Role.SupportUser]
						},
					},
				],
				meta: {
					roles: [Role.SupportUser]
				},
			},
			{
				path: "Profile",
				name: "ProfileSupportUser",
				component: Profile,
				meta: {
					roles: [Role.SupportUser]
				},
			},
		],
		meta: {
			roles: [Role.SupportUser]
		},
	},
	{
		path: "/Login",
		name: "Login",
		component: Login
	},
	{
		path: "/Logout",
		name: "Logout",
		component: Logout
	},
	{
		path: '*',
		redirect: "/Login"
	}
]

const router = new VueRouter({
	mode: 'history',
	base: process.env.BASE_URL,
	routes
})

export default router
router.beforeEach((to, from, next) => {
	const { roles } = to.meta;
	if (roles) {
		const userRole = JSON.parse(atob(sessionStorage.getItem('token').split('.')[1])).role;
		if(roles.length && !roles.includes(userRole)){
			return next({path: 'Login'});
		}
	}
	next();
});
