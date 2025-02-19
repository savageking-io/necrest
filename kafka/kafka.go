package kafka

import (
	"fmt"
	"github.com/IBM/sarama"
	log "github.com/sirupsen/logrus"
	log2 "log"
	"os"
)

type KafkaConfig struct {
	ClientID string `yaml:"client_id"`
	Host     string `yaml:"host"`
	Port     uint16 `yaml:"port"`
}

type Kafka struct {
	Config   *sarama.Config
	Producer sarama.SyncProducer
}

func (d *Kafka) Init(config *KafkaConfig) error {
	log.Traceln("REST::Init")

	if config.ClientID == "" {
		config.ClientID = "nec"
	}
	if config.Host == "" {
		config.Host = "localhost"
	}
	if config.Port == 0 {
		config.Port = 9092
	}

	log.Infof("Initializing Kafka client %s: %s:%d", config.ClientID, config.Host, config.Port)

	var err error
	d.Config = sarama.NewConfig()
	d.Config.Producer.RequiredAcks = sarama.WaitForAll
	d.Config.Producer.Retry.Max = 5
	d.Config.Producer.Return.Successes = true
	d.Config.ClientID = config.ClientID

	if log.GetLevel() == log.DebugLevel || log.GetLevel() == log.TraceLevel {
		sarama.Logger = log2.New(os.Stdout, "[SARAMA] ", log2.LstdFlags)
	}

	d.Producer, err = sarama.NewSyncProducer([]string{fmt.Sprintf("%s:%d", config.Host, config.Port)}, d.Config)
	if err != nil {
		return fmt.Errorf("failed to start Sarama producer: %s", err)
	}

	return nil
}

func (d *Kafka) SendMessage(topic string, message []byte) error {
	log.Tracef("REST::SendMessage. Topic: %s. Message: %s", topic, string(message))
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(message),
	}

	_, _, err := d.Producer.SendMessage(msg)
	return err
}
