// Code generated by goctl. DO NOT EDIT.
package types

type ServerPingReq struct {
	HttpCode int `form:"httpCode, default=200"`
}

type ServerPingRep struct {
	Result string `json:"result"`
}

type ServerPingDbReq struct {
	AsResult string `form:"asResult, default=carservice"`
}

type ServerPingDbRep struct {
	Result string `json:"result"`
}

type GetUserPhoneNumberReq struct {
	Code string `form:"code"`
}

type GetUserPhoneNumberRep struct {
	PhoneNumber string `json:"phoneNumber"`
}

type UploadFileRep struct {
	AbsolutePath string `json:"absolutePath"`
	RelativePath string `json:"relativePath"`
}

type UploadMultipleFilesRep struct {
	AbsolutePaths []string `json:"absolutePaths"`
	RelativePaths []string `json:"relativePaths"`
}

type UploadImageReq struct {
}

type UploadImageRep struct {
	Url string `json:"url"`
}

type WebsocketTestReq struct {
	Arg string `form:"arg"`
}

type WebsocketTestRep struct {
}

type WebsocketServiceListItem struct {
	ServiceId   uint8  `json:"serviceId"`
	ServiceDesc string `json:"serviceDesc"`
}

type WebsocketServiceReq struct {
}

type WebsocketServiceRep struct {
}

type SendCaptchaReq struct {
	PhoneNumber string `form:"phoneNumber"`
}

type SendCaptchaRep struct {
}

type PhoneNumberLoginReq struct {
	PhoneNumber string `json:"phoneNumber"`
	Captcha     string `json:"captcha"`
}

type PhoneNumberLoginRep struct {
	Token string `json:"token"`
}

type GetUserByPhoneNumberReq struct {
	PhoneNumber string `form:"phoneNumber"`
}

type GetUserByPhoneNumberRep struct {
	Username  string `json:"username"`
	AvatarUrl string `json:"avatarUrl"`
}

type WechatAuthorizationReq struct {
	Code string `form:"code"`
}

type WechatAuthorizationRep struct {
	Token string `json:"token"`
}

type MockLoginReq struct {
	Token string `json:"token"`
}

type GetUserProfileReq struct {
}

type GetUserProfileRep struct {
	Id          uint   `json:"id"`
	Username    string `json:"username"`
	PhoneNumber string `json:"phoneNumber"`
	AvatarUrl   string `json:"avatarUrl"`
}

type UpdateUserProfileReq struct {
	Username  string `json:"username"`
	AvatarUrl string `json:"avatarUrl"`
}

type UpdateUserProfileRep struct {
}

type CarBrandOptionListItem struct {
	Id     uint   `json:"id"`
	Label  string `json:"label"`
	Pinyin string `json:"pinyin"`
}

type CarBrandOptionListReq struct {
	Page int `form:"page,optional"`
}

type CarBrandOptionListRep struct {
}

type CarBrandSeriesOptionListItem struct {
	Id     uint   `json:"id"`
	Label  string `json:"label"`
	Pinyin string `json:"pinyin"`
}

type GetCarBrandSeriesOptionListReq struct {
	BrandId uint `form:"brandId"`
	Page    int  `form:"page,optional"`
}

type GetCarBrandSeriesOptionListRep struct {
}

type CarOwnerInfoCheckEmptyListRep struct {
	Listable bool `json:"listable"`
}

type CreateCarOwnerInfoReq struct {
	Name              string  `json:"name"`
	PhoneNumber       string  `json:"phoneNumber"`
	MultilevelAddress string  `json:"multilevelAddress"`
	FullAddress       string  `json:"fullAddress"`
	Longitude         float64 `json:"longitude"`
	Latitude          float64 `json:"latitude"`
}

type CreateCarOwnerInfoRep struct {
}

type UpdateCarOwnerInfoReq struct {
	Id                uint    `path:"id"`
	Name              string  `json:"name"`
	PhoneNumber       string  `json:"phoneNumber"`
	MultilevelAddress string  `json:"multilevelAddress"`
	FullAddress       string  `json:"fullAddress"`
	Longitude         float64 `json:"longitude"`
	Latitude          float64 `json:"latitude"`
}

type UpdateCarOwnerInfoRep struct {
}

type CarOwnerInfoListItem struct {
	Id                uint   `json:"id"`
	Name              string `json:"name"`
	PhoneNumber       string `json:"phoneNumber"`
	MultilevelAddress string `json:"multilevelAddress"`
	FullAddress       string `json:"fullAddress"`
}

type GetCarOwnerInfoListReq struct {
	Id uint `path:"id"`
}

type GetCarOwnerInfoListRep struct {
}

type GetCarOwnerInfoReq struct {
	Id uint `path:"id"`
}

type GetCarOwnerInfoRep struct {
	Id                uint   `json:"id"`
	Name              string `json:"name"`
	PhoneNumber       string `json:"phoneNumber"`
	MultilevelAddress string `json:"multilevelAddress"`
	FullAddress       string `json:"fullAddress"`
}

type DeleteCarOwnerInfoReq struct {
	Id uint `path:"id"`
}

type UserOrderListItem struct {
	Id           uint   `json:"id"`
	Deletable    bool   `json:"deletable"`
	OrderNumber  string `json:"orderNumber"`
	PartnerStore string `json:"partnerStore"`
	Requirements string `json:"requirements"`
	OrderStatus  string `json:"orderStatus"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
}

type CreateUserOrderReq struct {
	CarOwnerName        string  `json:"carOwnerName"`
	CarOwnerPhoneNumber string  `json:"carOwnerPhoneNumber"`
	CarOwnerLongitude   float64 `json:"carOwnerLongitude"`
	CarOwnerLatitude    float64 `json:"carOwnerLatitude"`
	CarOwnerMultiLvAddr string  `json:"carOwnerMultiLvAddr"`
	CarOwnerFullAddress string  `json:"carOwnerFullAddress"`
	CarBrandId          int64   `json:"carBrandId"`
	CarSeriesId         int64   `json:"carSeriesId"`
	PartnerStoreId      int64   `json:"partnerStoreId,optional"`
	Requirements        string  `json:"requirements"`
	AgreeToTerms        uint8   `json:"agreeToTerms"`
}

type CreateUserOrderRep struct {
	NewId uint `json:"newId"`
}

type UpdateUserOrderReq struct {
	CarOwnerName        string  `json:"carOwnerName"`
	CarOwnerPhoneNumber string  `json:"carOwnerPhoneNumber"`
	CarOwnerLongitude   float64 `json:"carOwnerLongitude"`
	CarOwnerLatitude    float64 `json:"carOwnerLatitude"`
	CarOwnerMultLvAddr  string  `json:"carOwnerMultLvAddr"`
	CarOwnerFullAddress string  `json:"carOwnerFullAddress"`
	CarBrandId          uint    `json:"carBrandId"`
	CarSeriesId         uint    `json:"carSeriesId"`
	PartnerStoreId      uint    `json:"partnerStoreId"`
	Requirements        string  `json:"requirements"`
}

type UpdateUserOrderRep struct {
}

type GetUserOrderReq struct {
	Id uint `path:"id"`
}

type GetUserOrderRep struct {
	Id                  uint   `json:"id"`
	OrderNumber         string `json:"orderNumber"`
	CarOwnerName        string `json:"carOwnerName"`
	CarOwnerMultiLvAddr string `json:"carOwnerMultiLvAddr"`
	CarOwnerFullAddr    string `json:"carOwnerFullAddr"`
	CarBrandName        string `json:"carBrandName"`
	CarSeriesName       string `json:"carSeriesName"`
	PartnerStore        string `json:"partnerStore"`
	PartnerStoreAddr    string `json:"partnerStoreAddr"`
	Requirements        string `json:"requirements"`
	OrderStatus         string `json:"orderStatus"`
	CreatedAt           string `json:"createdAt"`
	UpdatedAt           string `json:"updatedAt"`
}

type GetUserOrderListReq struct {
	Page int `form:"page,optional"`
}

type GetUserOrderListRep struct {
}

type DeleteUserOrderReq struct {
	Id uint `path:"id"`
}

type PartnerStoreListItem struct {
	Id          uint    `json:"id"`
	Title       string  `json:"title"`
	FullAddress string  `json:"fullAddress"`
	Gap         float32 `json:"gap"`
	Unit        string  `json:"unit"`
}

type GetPartnerStoreListReq struct {
	Address  string  `form:"address"`
	LimitGap float32 `form:"limitGap,optional"`
}

type GetPartnerStoreListRep struct {
}
