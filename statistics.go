// https://github.com/Haivision/srt/blob/master/docs/API/statistics.md

package srt

// Statistics represents the statistics for a connection
type Statistics struct {
	MsTimeStamp uint64 `json:"ms_time_stamp"` // The time elapsed, in milliseconds, since the SRT socket has been created

	// Accumulated
	Accumulated StatisticsAccumulated `json:"accumulated"`

	// Interval
	Interval StatisticsInterval `json:"interval"`

	// Instantaneous
	Instantaneous StatisticsInstantaneous `json:"instantaneous"`
}

type StatisticsAccumulated struct {
	PktSent          uint64 `json:"pkt_sent"`         // The total number of sent DATA packets, including retransmitted packets
	PktRecv          uint64 `json:"pkt_recv"`         // The total number of received DATA packets, including retransmitted packets
	PktSentUnique    uint64 `json:"pkt_sent_unique"`  // The total number of unique DATA packets sent by the SRT sender
	PktRecvUnique    uint64 `json:"pkt_recv_unique"`  // The total number of unique original, retransmitted or recovered by the packet filter DATA packets received in time, decrypted without errors and, as a result, scheduled for delivery to the upstream application by the SRT receiver.
	PktSendLoss      uint64 `json:"pkt_send_loss"`    // The total number of data packets considered or reported as lost at the sender side. Does not correspond to the packets detected as lost at the receiver side.
	PktRecvLoss      uint64 `json:"pkt_recv_loss"`    // The total number of SRT DATA packets detected as presently missing (either reordered or lost) at the receiver side
	PktRetrans       uint64 `json:"pkt_retrans"`      // The total number of retransmitted packets sent by the SRT sender
	PktRecvRetrans   uint64 `json:"pkt_recv_retrans"` // The total number of retransmitted packets registered at the receiver side
	PktSentACK       uint64 `json:"pkt_sent_ack"`     // The total number of sent ACK (Acknowledgement) control packets
	PktRecvACK       uint64 `json:"pkt_recv_ack"`     // The total number of received ACK (Acknowledgement) control packets
	PktSentNAK       uint64 `json:"pkt_sent_nak"`     // The total number of sent NAK (Negative Acknowledgement) control packets
	PktRecvNAK       uint64 `json:"pkt_recv_nak"`     // The total number of received NAK (Negative Acknowledgement) control packets
	PktSentKM        uint64 `json:"pkt_sent_km"`      // The total number of sent KM (Key Material) control packets
	PktRecvKM        uint64 `json:"pkt_recv_km"`      // The total number of received KM (Key Material) control packets
	UsSndDuration    uint64 `json:"us_snd_duration"`  // The total accumulated time in microseconds, during which the SRT sender has some data to transmit, including packets that have been sent, but not yet acknowledged
	PktRecvBelated   uint64 `json:"pkt_recv_belated"`
	PktSendDrop      uint64 `json:"pkt_send_drop"`      // The total number of dropped by the SRT sender DATA packets that have no chance to be delivered in time
	PktRecvDrop      uint64 `json:"pkt_recv_drop"`      // The total number of dropped by the SRT receiver and, as a result, not delivered to the upstream application DATA packets
	PktRecvUndecrypt uint64 `json:"pkt_recv_undecrypt"` // The total number of packets that failed to be decrypted at the receiver side

	ByteSent          uint64 `json:"byte_sent"`         // Same as pktSent, but expressed in bytes, including payload and all the headers (IP, TCP, SRT)
	ByteRecv          uint64 `json:"byte_recv"`         // Same as pktRecv, but expressed in bytes, including payload and all the headers (IP, TCP, SRT)
	ByteSentUnique    uint64 `json:"byte_sent_unique"`  // Same as pktSentUnique, but expressed in bytes, including payload and all the headers (IP, TCP, SRT)
	ByteRecvUnique    uint64 `json:"byte_recv_unique"`  // Same as pktRecvUnique, but expressed in bytes, including payload and all the headers (IP, TCP, SRT)
	ByteRecvLoss      uint64 `json:"byte_recv_loss"`    // Same as pktRecvLoss, but expressed in bytes, including payload and all the headers (IP, TCP, SRT), bytes for the presently missing (either reordered or lost) packets' payloads are estimated based on the average packet size
	ByteRetrans       uint64 `json:"byte_retrans"`      // Same as pktRetrans, but expressed in bytes, including payload and all the headers (IP, TCP, SRT)
	ByteRecvRetrans   uint64 `json:"byte_recv_retrans"` // Same as pktRecvRetrans, but expressed in bytes, including payload and all the headers (IP, TCP, SRT)
	ByteRecvBelated   uint64 `json:"byte_recv_belated"`
	ByteSendDrop      uint64 `json:"byte_send_drop"`      // Same as pktSendDrop, but expressed in bytes, including payload and all the headers (IP, TCP, SRT)
	ByteRecvDrop      uint64 `json:"byte_recv_drop"`      // Same as pktRecvDrop, but expressed in bytes, including payload and all the headers (IP, TCP, SRT)
	ByteRecvUndecrypt uint64 `json:"byte_recv_undecrypt"` // Same as pktRecvUndecrypt, but expressed in bytes, including payload and all the headers (IP, TCP, SRT)

	TickerOverslept   uint64 `json:"ticker_overslept"`    // The number of times the SRT connection's internal ticker has overslept, i.e. the time when the next packet should have been sent has already passed at the moment when the sender is notified about the tick.
	MsTickerOverslept uint64 `json:"ms_ticker_overslept"` // The accumulated time in milliseconds the SRT connection's internal ticker has overslept.
}

type StatisticsInterval struct {
	MsInterval uint64 `json:"ms_interval"` // Length of the interval, in milliseconds

	PktSent        uint64 `json:"pkt_sent"`         // Number of sent DATA packets, including retransmitted packets
	PktRecv        uint64 `json:"pkt_recv"`         // Number of received DATA packets, including retransmitted packets
	PktSentUnique  uint64 `json:"pkt_sent_unique"`  // Number of unique DATA packets sent by the SRT sender
	PktRecvUnique  uint64 `json:"pkt_recv_unique"`  // Number of unique original, retransmitted or recovered by the packet filter DATA packets received in time, decrypted without errors and, as a result, scheduled for delivery to the upstream application by the SRT receiver.
	PktSendLoss    uint64 `json:"pkt_send_loss"`    // Number of data packets considered or reported as lost at the sender side. Does not correspond to the packets detected as lost at the receiver side.
	PktRecvLoss    uint64 `json:"pkt_recv_loss"`    // Number of SRT DATA packets detected as presently missing (either reordered or lost) at the receiver side
	PktRetrans     uint64 `json:"pkt_retrans"`      // Number of retransmitted packets sent by the SRT sender
	PktRecvRetrans uint64 `json:"pkt_recv_retrans"` // Number of retransmitted packets registered at the receiver side
	PktSentACK     uint64 `json:"pkt_sent_ack"`     // Number of sent ACK (Acknowledgement) control packets
	PktRecvACK     uint64 `json:"pkt_recv_ack"`     // Number of received ACK (Acknowledgement) control packets
	PktSentNAK     uint64 `json:"pkt_sent_nak"`     // Number of sent NAK (Negative Acknowledgement) control packets
	PktRecvNAK     uint64 `json:"pkt_recv_nak"`     // Number of received NAK (Negative Acknowledgement) control packets

	MbpsSendRate float64 `json:"mbps_send_rate"` // Sending rate, in Mbps
	MbpsRecvRate float64 `json:"mbps_recv_rate"` // Receiving rate, in Mbps

	UsSndDuration uint64 `json:"us_snd_duration"` // Accumulated time in microseconds, during which the SRT sender has some data to transmit, including packets that have been sent, but not yet acknowledged

	PktReorderDistance uint64 `json:"pkt_reorder_distance"`
	PktRecvBelated     uint64 `json:"pkt_recv_belated"`   // Number of packets that arrive too late
	PktSndDrop         uint64 `json:"pkt_snd_drop"`       // Number of dropped by the SRT sender DATA packets that have no chance to be delivered in time
	PktRecvDrop        uint64 `json:"pkt_recv_drop"`      // Number of dropped by the SRT receiver and, as a result, not delivered to the upstream application DATA packets
	PktRecvUndecrypt   uint64 `json:"pkt_recv_undecrypt"` // Number of packets that failed to be decrypted at the receiver side

	ByteSent          uint64 `json:"byte_sent"`           // Same as pktSent, but expressed in bytes, including payload and all the headers (IP, TCP, SRT)
	ByteRecv          uint64 `json:"byte_recv"`           // Same as pktRecv, but expressed in bytes, including payload and all the headers (IP, TCP, SRT)
	ByteSentUnique    uint64 `json:"byte_sent_unique"`    // Same as pktSentUnique, but expressed in bytes, including payload and all the headers (IP, TCP, SRT)
	ByteRecvUnique    uint64 `json:"byte_recv_unique"`    // Same as pktRecvUnique, but expressed in bytes, including payload and all the headers (IP, TCP, SRT)
	ByteRecvLoss      uint64 `json:"byte_recv_loss"`      // Same as pktRecvLoss, but expressed in bytes, including payload and all the headers (IP, TCP, SRT), bytes for the presently missing (either reordered or lost) packets' payloads are estimated based on the average packet size
	ByteRetrans       uint64 `json:"byte_retrans"`        // Same as pktRetrans, but expressed in bytes, including payload and all the headers (IP, TCP, SRT)
	ByteRecvRetrans   uint64 `json:"byte_recv_retrans"`   // Same as pktRecvRetrans, but expressed in bytes, including payload and all the headers (IP, TCP, SRT)
	ByteRecvBelated   uint64 `json:"byte_recv_belated"`   // Same as pktRecvBelated, but expressed in bytes, including payload and all the headers (IP, TCP, SRT)
	ByteSendDrop      uint64 `json:"byte_send_drop"`      // Same as pktSendDrop, but expressed in bytes, including payload and all the headers (IP, TCP, SRT)
	ByteRecvDrop      uint64 `json:"byte_recv_drop"`      // Same as pktRecvDrop, but expressed in bytes, including payload and all the headers (IP, TCP, SRT)
	ByteRecvUndecrypt uint64 `json:"byte_recv_undecrypt"` // Same as pktRecvUndecrypt, but expressed in bytes, including payload and all the headers (IP, TCP, SRT)

	TickerOverslept   uint64 `json:"ticker_overslept"`    // The number of times the SRT connection's internal ticker has overslept, i.e. the time when the next packet should have been sent has already passed at the moment when the sender is notified about the tick.
	MsTickerOverslept uint64 `json:"ms_ticker_overslept"` // The accumulated time in milliseconds the SRT connection's internal ticker has overslept.
}

type StatisticsInstantaneous struct {
	UsPktSendPeriod       float64 `json:"us_pkt_send_period"`        // Current minimum time interval between which consecutive packets are sent, in microseconds
	PktFlowWindow         uint64  `json:"pkt_flow_window"`           // The maximum number of packets that can be "in flight"
	PktFlightSize         uint64  `json:"pkt_flight_size"`           // The number of packets in flight
	MsRTT                 float64 `json:"ms_rtt"`                    // Smoothed round-trip time (SRTT), an exponentially-weighted moving average (EWMA) of an endpoint's RTT samples, in milliseconds
	MbpsSentRate          float64 `json:"mbps_sent_rate"`            // Current transmission bandwidth, in Mbps
	MbpsRecvRate          float64 `json:"mbps_recv_rate"`            // Current receiving bandwidth, in Mbps
	MbpsLinkCapacity      float64 `json:"mbps_link_capacity"`        // Estimated capacity of the network link, in Mbps
	ByteAvailSendBuf      uint64  `json:"byte_avail_send_buf"`       // The available space in the sender's buffer, in bytes
	ByteAvailRecvBuf      uint64  `json:"byte_avail_recv_buf"`       // The available space in the receiver's buffer, in bytes
	MbpsMaxBW             float64 `json:"mbps_max_bw"`               // Transmission bandwidth limit, in Mbps
	ByteMSS               uint64  `json:"byte_mss"`                  // Maximum Segment Size (MSS), in bytes
	PktSendBuf            uint64  `json:"pkt_send_buf"`              // The number of packets in the sender's buffer that are already scheduled for sending or even possibly sent, but not yet acknowledged
	ByteSendBuf           uint64  `json:"byte_send_buf"`             // Instantaneous (current) value of pktSndBuf, but expressed in bytes, including payload and all headers (IP, TCP, SRT)
	MsSendBuf             uint64  `json:"ms_send_buf"`               // The timespan (msec) of packets in the sender's buffer (unacknowledged packets)
	MsSendTsbPdDelay      uint64  `json:"ms_send_tsb_pd_delay"`      // Timestamp-based Packet Delivery Delay value of the peer
	PktRecvBuf            uint64  `json:"pkt_recv_buf"`              // The number of acknowledged packets in receiver's buffer
	ByteRecvBuf           uint64  `json:"byte_recv_buf"`             // Instantaneous (current) value of pktRcvBuf, expressed in bytes, including payload and all headers (IP, TCP, SRT)
	MsRecvBuf             uint64  `json:"ms_recv_buf"`               // The timespan (msec) of acknowledged packets in the receiver's buffer
	MsRecvTsbPdDelay      uint64  `json:"ms_recv_tsb_pd_delay"`      // Timestamp-based Packet Delivery Delay value set on the socket via SRTO_RCVLATENCY or SRTO_LATENCY
	PktReorderTolerance   uint64  `json:"pkt_reorder_tolerance"`     // Instant value of the packet reorder tolerance
	PktRecvAvgBelatedTime uint64  `json:"pkt_recv_avg_belated_time"` // Accumulated difference between the current time and the time-to-play of a packet that is received late
	PktSendLossRate       float64 `json:"pkt_send_loss_rate"`        // Percentage of resent data vs. sent data
	PktRecvLossRate       float64 `json:"pkt_recv_loss_rate"`        // Percentage of retransmitted data vs. received data
}
